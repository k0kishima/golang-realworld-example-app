package validators_test

import (
	"testing"

	"github.com/k0kishima/golang-realworld-example-app/ent"
	"github.com/k0kishima/golang-realworld-example-app/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidateArticle(t *testing.T) {
	tests := []struct {
		name     string
		article  *ent.Article
		expected validators.ArticleValidationResult
	}{
		{
			name: "valid article",
			article: &ent.Article{
				Title:       "Test Article",
				Description: "Test Description",
				Body:        "Test Body",
			},
			expected: validators.ArticleValidationResult{
				Valid:  true,
				Errors: nil,
			},
		},
		{
			name: "missing title",
			article: &ent.Article{
				Description: "Test Description",
				Body:        "Test Body",
			},
			expected: validators.ArticleValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"title": {"can't be blank"},
				},
			},
		},
		{
			name: "missing description",
			article: &ent.Article{
				Title: "Test Article",
				Body:  "Test Body",
			},
			expected: validators.ArticleValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"description": {"can't be blank"},
				},
			},
		},
		{
			name: "missing body",
			article: &ent.Article{
				Title:       "Test Article",
				Description: "Test Description",
			},
			expected: validators.ArticleValidationResult{
				Valid: false,
				Errors: map[string][]string{
					"body": {"can't be blank"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidateArticle(tt.article)
			assert.Equal(t, tt.expected, result)
		})
	}
}
