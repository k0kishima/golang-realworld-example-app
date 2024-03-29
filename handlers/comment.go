package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/article"
	"github.com/k0kishima/golang-realworld-example-app/ent/comment"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
)

func GetComments(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		targetArticle, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		comments, err := targetArticle.QueryComments().All(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching comments"})
			return
		}

		// OPTIMIZE: Fix n + 1
		var commentsResponse []gin.H
		for _, comment := range comments {
			author, err := client.User.Query().Where(user.IDEQ(comment.AuthorID)).Only(c.Request.Context())
			if err != nil {
				respondWithError(c, http.StatusInternalServerError, "Error fetching comment author")
				return
			}

			currentUser, ok := GetCurrentUserFromContext(c)
			following := false
			if ok {
				isFollowing, err := isFollowing(c, currentUser, author)
				if err != nil {
					respondWithError(c, http.StatusInternalServerError, "Error checking if user is following")
				}
				following = isFollowing
			}

			commentResponse := gin.H{
				"id":        comment.ID,
				"body":      comment.Body,
				"createdAt": formatTimeForAPI(comment.CreatedAt),
				"updatedAt": formatTimeForAPI(comment.UpdatedAt),
				"author": gin.H{
					"username":  author.Username,
					"bio":       author.Bio,
					"image":     author.Image,
					"following": following,
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

		currentUser, ok := GetCurrentUserFromContext(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		slug := c.Param("slug")
		targetArticle, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		following, err := currentUser.QueryFollowing().Where(user.IDEQ(targetArticle.AuthorID)).Exist(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking if user is following"})
			return
		}

		comment, err := client.Comment.Create().
			SetBody(req.Comment.Body).
			SetAuthorID(currentUser.ID).
			Save(c.Request.Context())
		targetArticle.Update().AddComments(comment).Save(c.Request.Context())

		if err != nil {
			log.Printf("Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"comment": gin.H{
				"id":        comment.ID,
				"body":      comment.Body,
				"createdAt": formatTimeForAPI(comment.CreatedAt),
				"updatedAt": formatTimeForAPI(comment.UpdatedAt),
				"author": gin.H{
					"username":  currentUser.Username,
					"bio":       currentUser.Bio,
					"image":     currentUser.Image,
					"following": following,
				},
			},
		})
	}
}

func DeleteComment(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := GetCurrentUserFromContext(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error asserting user type"})
			return
		}

		slug := c.Param("slug")
		_, err := client.Article.Query().Where(article.SlugEQ(slug)).Only(c.Request.Context())
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		commentID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid comment ID"})
			return
		}
		targetComment, err := client.Comment.Query().Where(comment.IDEQ(commentID)).Only(c.Request.Context())
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		if targetComment.AuthorID != currentUser.ID {
			c.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this comment"})
			return
		}

		err = client.Comment.DeleteOne(targetComment).Exec(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting comment"})
			return
		}

		c.Status(http.StatusOK)
	}
}
