package validators

import (
	"context"

	"github.com/k0kishima/golang-realworld-example-app/ent"
	entarticle "github.com/k0kishima/golang-realworld-example-app/ent/article"
)

type ArticleValidationResult struct {
	Valid  bool
	Errors map[string][]string
}

func ValidateArticle(client *ent.Client, article *struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tagList"`
}) ArticleValidationResult {
	errors := make(map[string][]string)

	if article.Title == "" {
		errors["title"] = append(errors["title"], "can't be blank")
	} else {
		exists, err := client.Article.Query().Where(entarticle.TitleEQ(article.Title)).Exist(context.Background())
		if err != nil {
			errors["title"] = append(errors["title"], "error checking title uniqueness")
		} else if exists {
			errors["title"] = append(errors["title"], "must be unique")
		}
	}

	if article.Description == "" {
		errors["description"] = append(errors["description"], "can't be blank")
	}
	if article.Body == "" {
		errors["body"] = append(errors["body"], "can't be blank")
	}

	return ArticleValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}
