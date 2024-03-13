package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/auth"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfollow"
)

func GetProfile(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		targetUser, err := client.User.Query().Where(user.UsernameEQ(username)).Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}

		following := false
		token := c.GetHeader("Authorization")
		if token != "" {
			claims, err := auth.ParseToken(token)
			if err == nil {
				currentUser, err := client.User.Query().Where(user.EmailEQ(claims.Email)).Only(c.Request.Context())
				if err == nil {
					exists, err := currentUser.QueryFollows().Where(userfollow.FolloweeIDEQ(targetUser.ID)).Exist(c.Request.Context())
					if err == nil && exists {
						following = true
					}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"profile": gin.H{
				"username":  targetUser.Username,
				"bio":       targetUser.Bio,
				"image":     targetUser.Image,
				"following": following,
			},
		})
	}
}
