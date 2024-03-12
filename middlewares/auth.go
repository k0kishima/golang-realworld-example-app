package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/auth"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
)

func AuthMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "missing authorization credentials"})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		u, err := client.User.Query().Where(user.EmailEQ(claims.Email)).Only(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching user"})
			c.Abort()
			return
		}

		c.Set("user", u)
		c.Next()
	}
}
