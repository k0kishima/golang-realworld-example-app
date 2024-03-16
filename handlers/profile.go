package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfollow"
)

func GetProfile(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		targetUser, err := getUserByUsername(client, c, username)
		if err != nil {
			if ent.IsNotFound(err) {
				respondWithError(c, http.StatusNotFound, "Profile not found")
			} else {
				respondWithError(c, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		following := false
		currentUser, _ := c.Get("currentUser")
		currentUserEntity, exists := currentUser.(*ent.User)
		if exists {
			exists, err := currentUserEntity.QueryFollows().Where(userfollow.FolloweeIDEQ(targetUser.ID)).Exist(c.Request.Context())
			if err == nil && exists {
				following = true
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
			return
		}

		if err := followUser(client, c, currentUserEntity, targetUser); err != nil {
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

		if err := unfollowUser(client, c, currentUserEntity, targetUser); err != nil {
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

func followUser(client *ent.Client, c *gin.Context, currentUserEntity, targetUser *ent.User) error {
	_, err := client.UserFollow.Create().
		SetFollower(currentUserEntity).
		SetFollowee(targetUser).
		Save(c.Request.Context())
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

func unfollowUser(client *ent.Client, c *gin.Context, currentUserEntity, targetUser *ent.User) error {
	_, err := client.UserFollow.Delete().Where(
		userfollow.And(
			userfollow.FollowerIDEQ(currentUserEntity.ID),
			userfollow.FolloweeIDEQ(targetUser.ID),
		),
	).Exec(c.Request.Context())
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error unfollowing user")
		return err
	}
	return nil
}
