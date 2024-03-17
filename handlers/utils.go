package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func handleCommonErrors(c *gin.Context, err error) {
	if ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	}
}

func formatTimeForAPI(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

func GetCurrentUserFromContext(c *gin.Context) (*ent.User, bool) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		return nil, false
	}
	currentUserEntity, ok := currentUser.(*ent.User)
	return currentUserEntity, ok
}

func getUserByUsername(client *ent.Client, c *gin.Context, username string) (*ent.User, error) {
	return client.User.Query().Where(user.UsernameEQ(username)).Only(c.Request.Context())
}

func isFollowing(c *gin.Context, follower *ent.User, followee *ent.User) (bool, error) {
	exists, err := follower.QueryFollowing().Where(user.IDEQ(followee.ID)).Exist(c.Request.Context())
	if err != nil {
		return false, err
	}
	return exists, nil
}
