package validators

import "github.com/k0kishima/golang-realworld-example-app/ent"

type ArticleValidationResult struct {
	Valid  bool
	Errors map[string][]string
}

func ValidateArticle(article *ent.Article) ArticleValidationResult {
	errors := make(map[string][]string)

	if article.Title == "" {
		errors["title"] = append(errors["title"], "can't be blank")
	}
	if article.Description == "" {
		errors["description"] = append(errors["description"], "can't be blank")
	}
	if article.Body == "" {
		errors["body"] = append(errors["body"], "can't be blank")
	}

	if len(errors) == 0 {
		return ArticleValidationResult{
			Valid:  true,
			Errors: nil,
		}
	}

	return ArticleValidationResult{
		Valid:  false,
		Errors: errors,
	}
}
