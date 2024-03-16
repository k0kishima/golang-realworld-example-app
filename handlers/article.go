package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/auth"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/article"
	"github.com/k0kishima/golang-realworld-example-app/ent/articletag"
	"github.com/k0kishima/golang-realworld-example-app/ent/comment"
	"github.com/k0kishima/golang-realworld-example-app/ent/tag"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfavorite"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfollow"
	"github.com/k0kishima/golang-realworld-example-app/validators"
)

func GetArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		tagList, err := article.QueryTags().Select(tag.FieldDescription).Strings(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tags"})
			return
		}

		currentUser, err := getCurrentUser(client, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		favorited, favoritesCount, err := getArticleFavoritedAndCount(client, article, currentUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites information"})
			return
		}

		response, err := articleResponse(client, article, tagList, favorited, favoritesCount, currentUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"article": response})
	}
}

func ListArticles(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := client.Article.Query().WithArticleAuthor().WithTags()

		if tagName := c.Query("tag"); tagName != "" {
			query.Where(article.HasTagsWith(tag.DescriptionEQ(tagName)))
		}

		if author := c.Query("author"); author != "" {
			query.Where(article.HasArticleAuthorWith(user.UsernameEQ(author)))
		}

		if favorited := c.Query("favorited"); favorited != "" {
			query.Where(article.HasFavoritedUsersWith(user.UsernameEQ(favorited)))
		}

		limitStr := c.DefaultQuery("limit", "20")
		offsetStr := c.DefaultQuery("offset", "0")

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit value"})
			return
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid offset value"})
			return
		}

		articles, err := query.Limit(limit).Offset(offset).All(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching articles"})
			return
		}

		currentUser, _ := getCurrentUser(client, c)

		articlesResponse := []gin.H{}
		for _, article := range articles {
			tagList := make([]string, len(article.Edges.Tags))
			for i, tag := range article.Edges.Tags {
				tagList[i] = tag.Description
			}

			favorited, favoritesCount, err := getArticleFavoritedAndCount(client, article, currentUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites information"})
				return
			}

			response, err := articleResponse(client, article, tagList, favorited, favoritesCount, currentUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
				return
			}
			articlesResponse = append(articlesResponse, response)
		}

		c.JSON(http.StatusOK, gin.H{
			"articles":      articlesResponse,
			"articlesCount": len(articlesResponse),
		})
	}
}

func CreateArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Article struct {
				Title       string   `json:"title"`
				Description string   `json:"description"`
				Body        string   `json:"body"`
				TagList     []string `json:"tagList"`
			} `json:"article"`
		}
		if err := c.BindJSON(&req); err != nil {
			respondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		validationResult := validators.ValidateArticle(&ent.Article{
			Title:       req.Article.Title,
			Description: req.Article.Description,
			Body:        req.Article.Body,
		})
		if !validationResult.Valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationResult.Errors})
			return
		}

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		tagIDs, err := findOrCreateTagIDsByNames(client, req.Article.TagList)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error processing tags"})
			return
		}

		tx, err := client.Tx(c.Request.Context())
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Something went wrong")
			return
		}

		article, err := tx.Article.Create().
			SetArticleAuthor(currentUserEntity).
			SetSlug(req.Article.Title).
			SetTitle(req.Article.Title).
			SetDescription(req.Article.Description).
			SetBody(req.Article.Body).
			Save(c.Request.Context())
		if err != nil {
			tx.Rollback()
			if ent.IsConstraintError(err) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": gin.H{"title": []string{"must be unique"}}})
			} else {
				respondWithError(c, http.StatusInternalServerError, "Error creating article")
			}
			return
		}

		var articleTagBuilders []*ent.ArticleTagCreate
		for _, tagID := range tagIDs {
			builder := tx.ArticleTag.Create().
				SetArticle(article).
				SetTagID(tagID)
			articleTagBuilders = append(articleTagBuilders, builder)
		}

		_, err = tx.ArticleTag.CreateBulk(articleTagBuilders...).Save(c.Request.Context())
		if err != nil {
			tx.Rollback()
			respondWithError(c, http.StatusInternalServerError, "Error creating article_tags")
			return
		}

		err = tx.Commit()
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Something went wrong")
		}

		favorited := false
		favoritesCount := 0
		response, err := articleResponse(client, article, req.Article.TagList, favorited, favoritesCount, currentUserEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"article": response})
	}
}

func UpdateArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Article struct {
				Title       string   `json:"title"`
				Description string   `json:"description"`
				Body        string   `json:"body"`
				TagList     []string `json:"tagList"`
			} `json:"article"`
		}
		if err := c.BindJSON(&req); err != nil {
			respondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		slug := c.Param("slug")
		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		update := client.Article.UpdateOneID(article.ID)
		if req.Article.Title != "" {
			update.SetSlug(req.Article.Title)
			update.SetTitle(req.Article.Title)
		}
		if req.Article.Description != "" {
			update.SetDescription(req.Article.Description)
		}
		if req.Article.Body != "" {
			update.SetBody(req.Article.Body)
		}

		updatedArticle, err := update.Save(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating article"})
			return
		}

		tagList, err := updatedArticle.QueryTags().Select(tag.FieldDescription).Strings(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tags"})
			return
		}

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		favorited, favoritesCount, err := getArticleFavoritedAndCount(client, updatedArticle, currentUserEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites information"})
			return
		}

		response, err := articleResponse(client, updatedArticle, tagList, favorited, favoritesCount, currentUserEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"article": response})
	}
}

func DeleteArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		article, err := client.Article.Query().Where(article.SlugEQ(slug)).WithArticleAuthor().Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		if article.Edges.ArticleAuthor.ID != currentUserEntity.ID {
			c.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this article"})
			return
		}

		tx, err := client.Tx(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error starting transaction"})
			return
		}

		_, err = tx.ArticleTag.Delete().Where(articletag.ArticleIDEQ(article.ID)).Exec(c.Request.Context())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting article tags"})
			return
		}

		_, err = tx.UserFavorite.Delete().Where(userfavorite.ArticleIDEQ(article.ID)).Exec(c.Request.Context())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting article favorites"})
			return
		}

		err = tx.Article.DeleteOne(article).Exec(c.Request.Context())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting article"})
			return
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error committing transaction"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func GetFeed(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		articles, err := client.Article.Query().
			Where(
				article.HasArticleAuthorWith(
					user.HasFollowsWith(
						userfollow.FolloweeIDEQ(currentUserEntity.ID),
					),
				),
			).
			Order(ent.Desc(article.FieldCreatedAt)).
			WithTags().
			All(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching feed"})
			return
		}

		articlesResponse := []gin.H{}
		for _, article := range articles {
			tagList := make([]string, len(article.Edges.Tags))
			for i, tag := range article.Edges.Tags {
				tagList[i] = tag.Description
			}

			favoritesCount, err := article.QueryFavoritedUsers().Count(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching favorites count"})
				return
			}

			// TODO: need to set favaorited
			response, err := articleResponse(client, article, tagList, false, favoritesCount, currentUserEntity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
				return
			}
			articlesResponse = append(articlesResponse, response)
		}

		c.JSON(http.StatusOK, gin.H{
			"articles":      articlesResponse,
			"articlesCount": len(articlesResponse),
		})
	}
}

func FavoriteArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		_, err = client.UserFavorite.Create().
			SetUser(currentUserEntity).
			SetArticle(article).
			Save(c.Request.Context())
		if err != nil {
			if ent.IsConstraintError(err) {
				c.JSON(http.StatusConflict, gin.H{"message": "Article is already favorited"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error favoriting article"})
			}
			return
		}

		tagList, err := article.QueryTags().Select(tag.FieldDescription).Strings(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tags"})
			return
		}

		favoritesCount, err := article.QueryFavoritedUsers().Count(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites count"})
			return
		}

		response, err := articleResponse(client, article, tagList, true, favoritesCount, currentUserEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"article": response})
	}
}

func UnfavoriteArticle(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		_, err = client.UserFavorite.Delete().
			Where(
				userfavorite.And(
					userfavorite.UserIDEQ(currentUserEntity.ID),
					userfavorite.ArticleIDEQ(article.ID),
				),
			).Exec(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error unfavoriting article"})
			return
		}

		tagList, err := article.QueryTags().Select(tag.FieldDescription).Strings(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tags"})
			return
		}

		favoritesCount, err := article.QueryFavoritedUsers().Count(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites count"})
			return
		}

		response, err := articleResponse(client, article, tagList, false, favoritesCount, currentUserEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating article response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"article": response})
	}
}

func GetComments(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		comments, err := article.QueryComments().WithCommentAuthor().All(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching comments"})
			return
		}

		var commentsResponse []gin.H
		for _, comment := range comments {
			commentResponse := gin.H{
				"id":        comment.ID,
				"body":      comment.Body,
				"createdAt": comment.CreatedAt,
				"updatedAt": comment.UpdatedAt,
				"author": gin.H{
					"username": comment.Edges.CommentAuthor.Username,
					"bio":      comment.Edges.CommentAuthor.Bio,
					"image":    comment.Edges.CommentAuthor.Image,
				},
			}
			commentsResponse = append(commentsResponse, commentResponse)
		}

		c.JSON(http.StatusOK, gin.H{"comments": commentsResponse})
	}
}

func PostComment(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Comment struct {
				Body string `json:"body"`
			} `json:"comment"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}

		if req.Comment.Body == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": gin.H{"body": []string{"can't be blank"}}})
			return
		}

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		slug := c.Param("slug")
		article, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		comment, err := client.Comment.Create().
			SetBody(req.Comment.Body).
			SetCommentAuthor(currentUserEntity).
			SetArticle(article).
			Save(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"comment": gin.H{
				"id":        comment.ID,
				"body":      comment.Body,
				"createdAt": comment.CreatedAt,
				"updatedAt": comment.UpdatedAt,
				"author": gin.H{
					"username": currentUserEntity.Username,
					"bio":      currentUserEntity.Bio,
					"image":    currentUserEntity.Image,
				},
			},
		})
	}
}

func DeleteComment(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error asserting user type"})
			return
		}

		slug := c.Param("slug")
		_, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		commentID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid comment ID"})
			return
		}
		comment, err := client.Comment.Query().Where(comment.IDEQ(commentID)).WithCommentAuthor().Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Comment not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching comment"})
			}
			return
		}

		if comment.Edges.CommentAuthor.ID != currentUserEntity.ID {
			c.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this comment"})
			return
		}

		err = client.Comment.DeleteOne(comment).Exec(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting comment"})
			return
		}

		c.Status(http.StatusOK)
	}
}

func findOrCreateTagIDsByNames(client *ent.Client, tagNames []string) ([]uuid.UUID, error) {
	var tagIDs []uuid.UUID
	for _, tagName := range tagNames {
		tag, err := client.Tag.Query().Where(tag.DescriptionEQ(tagName)).Only(context.Background())
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			tag, err = client.Tag.Create().SetDescription(tagName).Save(context.Background())
			if err != nil {
				return nil, err
			}
		}
		tagIDs = append(tagIDs, tag.ID)
	}
	return tagIDs, nil
}

func articleResponse(client *ent.Client, article *ent.Article, tagList []string, favorited bool, favoritesCount int, currentUser *ent.User) (gin.H, error) {
	author, err := article.QueryArticleAuthor().Only(context.Background())
	if err != nil {
		return nil, err
	}

	authorResponse, err := authorResponse(client, author, currentUser)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"slug":           article.Slug,
		"title":          article.Title,
		"description":    article.Description,
		"body":           article.Body,
		"tagList":        tagList,
		"favorited":      favorited,
		"favoritesCount": favoritesCount,
		"createdAt":      formatTimeForAPI(article.CreatedAt),
		"updatedAt":      formatTimeForAPI(article.UpdatedAt),
		"author":         authorResponse,
	}, nil
}

func authorResponse(client *ent.Client, author *ent.User, currentUser *ent.User) (gin.H, error) {
	following := false
	if currentUser != nil {
		isFollowing, err := client.UserFollow.Query().
			Where(userfollow.FollowerIDEQ(currentUser.ID), userfollow.FolloweeIDEQ(author.ID)).
			Exist(context.Background())
		if err != nil {
			return nil, err
		}
		following = isFollowing
	}

	return gin.H{
		"username":  author.Username,
		"bio":       author.Bio,
		"image":     author.Image,
		"following": following,
	}, nil
}

func getArticleFavoritedAndCount(client *ent.Client, article *ent.Article, currentUser *ent.User) (bool, int, error) {
	if currentUser == nil {
		return false, 0, nil
	}

	favorited, err := client.UserFavorite.Query().
		Where(
			userfavorite.And(
				userfavorite.UserIDEQ(currentUser.ID),
				userfavorite.ArticleIDEQ(article.ID),
			),
		).
		Exist(context.Background())
	if err != nil {
		return false, 0, err
	}

	favoritesCount, err := article.QueryFavoritedUsers().Count(context.Background())
	if err != nil {
		return false, 0, err
	}

	return favorited, favoritesCount, nil
}

func getCurrentUser(client *ent.Client, c *gin.Context) (*ent.User, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return nil, nil
	}

	claims, err := auth.ParseToken(token)
	if err != nil {
		return nil, err
	}

	return client.User.Query().Where(user.EmailEQ(claims.Email)).Only(c.Request.Context())
}
