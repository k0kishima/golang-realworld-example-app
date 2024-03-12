package validators_test

import (
	"testing"

	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserRegistration(t *testing.T) {
	tests := []struct {
		name     string
		user     *ent.User
		expected validators.UserValidationResult
	}{
		{
			name: "valid user",
			user: &ent.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: validators.UserValidationResult{
				Valid:  true,
				Errors: map[string][]string{},
			},
		},
		{
			name: "missing username",
			user: &ent.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: validators.UserValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"username": {"can't be blank"},
				},
			},
		},
		{
			name: "missing email",
			user: &ent.User{
				Username: "testuser",
				Password: "password123",
			},
			expected: validators.UserValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"email": {"can't be blank"},
				},
			},
		},
		{
			name: "missing password",
			user: &ent.User{
				Username: "testuser",
				Email:    "test@example.com",
			},
			expected: validators.UserValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"password": {"can't be blank"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidateUserRegistration(tt.user)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateUserLogin(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
		expected validators.UserValidationResult
	}{
		{
			name:     "valid credentials",
			email:    "test@example.com",
			password: "password123",
			expected: validators.UserValidationResult{
				Valid:  true,
				Errors: map[string][]string{},
			},
		},
		{
			name:     "missing email",
			password: "password123",
			expected: validators.UserValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"email": {"can't be blank"},
				},
			},
		},
		{
			name:  "missing password",
			email: "test@example.com",
			expected: validators.UserValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"password": {"can't be blank"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidateUserLogin(tt.email, tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}
