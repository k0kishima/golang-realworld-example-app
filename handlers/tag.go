package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/tag"
)

func GetTags(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tags, err := client.Tag.Query().Order(ent.Asc(tag.FieldDescription)).Limit(10).All(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tags"})
			return
		}

		var tagList []string
		for _, t := range tags {
			tagList = append(tagList, t.Description)
		}

		c.JSON(http.StatusOK, gin.H{"tags": tagList})
	}
}
