package validators

import (
	"github.com/k0kishima/golang-realworld-example-app/ent"
)

type UserValidationResult struct {
	Valid  bool
	Errors map[string][]string
}

func ValidateUserRegistration(user *ent.User) UserValidationResult {
	errors := make(map[string][]string)
	if user.Username == "" {
		errors["username"] = append(errors["username"], "can't be blank")
	}
	if user.Email == "" {
		errors["email"] = append(errors["email"], "can't be blank")
	}
	if user.Password == "" {
		errors["password"] = append(errors["password"], "can't be blank")
	}
	return UserValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}

func ValidateUserLogin(email, password string) UserValidationResult {
	errors := make(map[string][]string)
	if email == "" {
		errors["email"] = append(errors["email"], "can't be blank")
	}
	if password == "" {
		errors["password"] = append(errors["password"], "can't be blank")
	}
	return UserValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}
