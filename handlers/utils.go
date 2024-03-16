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

func formatTimeForAPI(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

func getUserByUsername(client *ent.Client, c *gin.Context, username string) (*ent.User, error) {
	targetUser, err := client.User.Query().Where(user.UsernameEQ(username)).Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			respondWithError(c, http.StatusNotFound, "User not found")
		} else {
			respondWithError(c, http.StatusInternalServerError, "Internal server error")
		}
		return nil, err
	}
	return targetUser, nil
}

func isFollowing(c *gin.Context, follower *ent.User, followee *ent.User) (bool, error) {
	exists, err := follower.QueryFollowing().Where(user.IDEQ(followee.ID)).Exist(c.Request.Context())
	if err != nil {
		return false, err
	}
	return exists, nil
}
