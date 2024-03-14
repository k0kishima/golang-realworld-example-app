package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/tag"
	"github.com/k0kishima/golang-realworld-example-app/validators"
)

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

		c.JSON(http.StatusCreated, gin.H{
			"article": gin.H{
				"title":       article.Title,
				"description": article.Description,
				"body":        article.Body,
				"tagList":     req.Article.TagList,
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
