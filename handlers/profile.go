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

		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		following := false
		if ok {
			isFollowing, err := isFollowing(c, currentUserEntity, targetUser)
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
		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			respondWithError(c, http.StatusInternalServerError, "Error asserting user type")
			return
		}

		username := c.Param("username")
		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			handleCommonErrors(c, err)
			return
		}

		if err := followUser(c, currentUserEntity, targetUser); err != nil {
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
		currentUser, _ := c.Get("currentUser")
		currentUserEntity, ok := currentUser.(*ent.User)
		if !ok {
			respondWithError(c, http.StatusInternalServerError, "Error asserting user type")
			return
		}

		username := c.Param("username")
		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			return
		}

		if err := unfollowUser(c, currentUserEntity, targetUser); err != nil {
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

func followUser(c *gin.Context, currentUserEntity, targetUser *ent.User) error {
	err := currentUserEntity.Update().AddFollowing(targetUser).Exec(c.Request.Context())
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

func unfollowUser(c *gin.Context, currentUserEntity, targetUser *ent.User) error {
	err := currentUserEntity.Update().RemoveFollowing(targetUser).Exec(c.Request.Context())
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error unfollowing user")
		return err
	}
	return nil
}
