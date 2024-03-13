package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/auth"
	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/utils"
	"github.com/k0kishima/golang-realworld-example-app/validators"
)

func RegisterUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			User struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Password string `json:"password"`
			} `json:"user"`
		}
		if err := c.BindJSON(&req); err != nil {
			respondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		validationResult := validators.ValidateUserRegistration(&ent.User{
			Username: req.User.Username,
			Email:    req.User.Email,
			Password: req.User.Password,
		})
		if !validationResult.Valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationResult.Errors})
			return
		}

		hashedPassword, err := utils.HashPassword(req.User.Password)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error hashing password")
			return
		}

		tx, err := client.Tx(c.Request.Context())
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error starting transaction")
			return
		}

		u, err := tx.User.
			Create().
			SetID(uuid.New()).
			SetUsername(req.User.Username).
			SetEmail(req.User.Email).
			SetPassword(hashedPassword).
			Save(c.Request.Context())

		if err != nil {
			tx.Rollback()
			handleUserCreationError(c, err)
			return
		}

		token, err := auth.CreateToken(u)
		if err != nil {
			tx.Rollback()
			respondWithError(c, http.StatusInternalServerError, "Error creating token")
			return
		}

		if err := tx.Commit(); err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error committing transaction")
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": gin.H{
			"username": u.Username,
			"email":    u.Email,
			"token":    token,
		}})
	}
}

func Login(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			User struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			} `json:"user"`
		}
		if err := c.BindJSON(&req); err != nil {
			respondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		validationResult := validators.ValidateUserLogin(req.User.Email, req.User.Password)
		if !validationResult.Valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationResult.Errors})
			return
		}

		u, err := client.User.Query().Where(user.EmailEQ(req.User.Email)).Only(c.Request.Context())
		if err != nil || !utils.CheckPasswordHash(req.User.Password, u.Password) {
			c.JSON(http.StatusForbidden, gin.H{"errors": gin.H{"email or password": []string{"is invalid"}}})
			return
		}

		token, err := auth.CreateToken(u)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error creating token")
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": gin.H{
			"username": u.Username,
			"email":    u.Email,
			"token":    token,
		}})
	}
}

func GetCurrentUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, _ := c.Get("currentUser")
		u := currentUser.(*ent.User)

		c.JSON(http.StatusOK, gin.H{"user": gin.H{
			"username": u.Username,
			"email":    u.Email,
		}})
	}
}

func UpdateUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			User struct {
				Username string `json:"username"`
				Email    string `json:"email"`
			} `json:"user"`
		}
		if err := c.BindJSON(&req); err != nil {
			respondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		currentUser, _ := c.Get("currentUser")
		u := currentUser.(*ent.User)

		u, err := u.Update().
			SetUsername(req.User.Username).
			SetEmail(req.User.Email).
			Save(c.Request.Context())

		if err != nil {
			handleUserUpdateError(c, err)
			return
		}

		token, err := auth.CreateToken(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": gin.H{
			"username": u.Username,
			"email":    u.Email,
			"token":    token,
		}})
	}
}

func handleUserCreationError(c *gin.Context, err error) {
	if ent.IsConstraintError(err) {
		errors := make(map[string][]string)
		if strings.Contains(err.Error(), "users.username") || strings.Contains(err.Error(), "users.email") {
			errors["username"] = []string{"has already been taken"}
			errors["email"] = []string{"has already been taken"}
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
	} else {
		respondWithError(c, http.StatusInternalServerError, "Error creating user")
	}
}

func handleUserUpdateError(c *gin.Context, err error) {
	if ent.IsConstraintError(err) {
		errors := make(map[string][]string)
		if strings.Contains(err.Error(), "users.username") {
			errors["username"] = []string{"has already been taken"}
		}
		if strings.Contains(err.Error(), "users.email") {
			errors["email"] = []string{"has already been taken"}
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
	} else {
		respondWithError(c, http.StatusInternalServerError, "Error updating user")
	}
}
