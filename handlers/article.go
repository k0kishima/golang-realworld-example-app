package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/auth"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/article"
	"github.com/k0kishima/golang-realworld-example-app/ent/articletag"
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

		favorited := false
		token := c.GetHeader("Authorization")
		if token != "" {
			claims, err := auth.ParseToken(token)
			if err == nil {
				currentUser, err := client.User.Query().Where(user.EmailEQ(claims.Email)).Only(c.Request.Context())
				if err == nil {
					favorited, err = isArticleFavoritedByUser(client, article, currentUser)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites"})
						return
					}
				}
			}
		}
		favoritesCount, err := client.UserFavorite.Query().
			Where(userfavorite.ArticleIDEQ(article.ID)).
			Count(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites count"})
			return
		}

		response := gin.H{
			"article": gin.H{
				"slug":           article.Slug,
				"title":          article.Title,
				"description":    article.Description,
				"body":           article.Body,
				"tagList":        tagList,
				"favorited":      favorited,
				"favoritesCount": favoritesCount,
			},
		}
		c.JSON(http.StatusOK, response)
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

		c.JSON(http.StatusCreated, articleResponse(article, req.Article.TagList))
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

		c.JSON(http.StatusOK, gin.H{
			"article": gin.H{
				"slug":        updatedArticle.Slug,
				"title":       updatedArticle.Title,
				"description": updatedArticle.Description,
				"body":        updatedArticle.Body,
				"tagList":     tagList,
			},
		})
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
			All(c.Request.Context())

		if err != nil && !ent.IsNotFound(err) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching feed"})
			return
		}

		articlesResponse := make([]gin.H, 0)
		for _, article := range articles {
			tagList, err := article.QueryTags().Select(tag.FieldDescription).Strings(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching tags"})
				return
			}

			articlesResponse = append(articlesResponse, articleResponse(article, tagList))
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

		favoritesCount, err := client.UserFavorite.Query().
			Where(userfavorite.ArticleIDEQ(article.ID)).
			Count(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching favorites count"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"article": gin.H{
				"slug":           article.Slug,
				"title":          article.Title,
				"description":    article.Description,
				"body":           article.Body,
				"tagList":        getTagList(article),
				"favorited":      true,
				"favoritesCount": favoritesCount,
			},
		})
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

func articleResponse(article *ent.Article, tagList []string) gin.H {
	return gin.H{
		"article": gin.H{
			"slug":        article.Slug,
			"title":       article.Title,
			"description": article.Description,
			"body":        article.Body,
			"tagList":     tagList,
		},
	}
}

func getTagList(article *ent.Article) []string {
	tags, err := article.QueryTags().All(context.Background())
	if err != nil {
		return nil
	}
	tagList := make([]string, len(tags))
	for i, tag := range tags {
		tagList[i] = tag.Description
	}
	return tagList
}

func isArticleFavoritedByUser(client *ent.Client, article *ent.Article, user *ent.User) (bool, error) {
	count, err := client.UserFavorite.Query().
		Where(userfavorite.And(
			userfavorite.UserIDEQ(user.ID),
			userfavorite.ArticleIDEQ(article.ID),
		)).
		Count(context.Background())
	return count > 0, err
}
