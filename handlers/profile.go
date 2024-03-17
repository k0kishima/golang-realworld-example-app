package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/ent"
)

func GetProfile(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		currentUser, ok := GetCurrentUserFromContext(c)
		following := false
		if ok {
			isFollowing, err := isFollowing(c, currentUser, targetUser)
			if err != nil {
				respondWithError(c, http.StatusInternalServerError, "Error checking if user is following")
			}
			following = isFollowing
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

func FollowUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := GetCurrentUserFromContext(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		username := c.Param("username")
		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		if err := followUser(c, currentUser, targetUser); err != nil {
			handleCommonErrors(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"profile": gin.H{
				"username":  targetUser.Username,
				"bio":       targetUser.Bio,
				"image":     targetUser.Image,
				"following": true,
			},
		})
	}
}

func UnfollowUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := GetCurrentUserFromContext(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error asserting user type"})
			return
		}

		username := c.Param("username")
		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			return
		}

		if err := unfollowUser(c, currentUser, targetUser); err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"profile": gin.H{
				"username":  targetUser.Username,
				"bio":       targetUser.Bio,
				"image":     targetUser.Image,
				"following": false,
			},
		})
	}
}

func followUser(c *gin.Context, currentUser, targetUser *ent.User) error {
	err := currentUser.Update().AddFollowing(targetUser).Exec(c.Request.Context())
	if err != nil {
		if ent.IsConstraintError(err) {
			respondWithError(c, http.StatusConflict, "User is already followed")
		} else {
			respondWithError(c, http.StatusInternalServerError, "Error following user")
		}
		return err
	}
	return nil
}

func unfollowUser(c *gin.Context, currentUser, targetUser *ent.User) error {
	err := currentUser.Update().RemoveFollowing(targetUser).Exec(c.Request.Context())
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error unfollowing user")
		return err
	}
	return nil
}
