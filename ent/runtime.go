// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/article"
	"github.com/k0kishima/golang-realworld-example-app/ent/schema"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfollow"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescSlug is the schema descriptor for slug field.
	articleDescSlug := articleFields[2].Descriptor()
	// article.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	article.SlugValidator = func() func(string) error {
		validators := articleDescSlug.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(slug string) error {
			for _, fn := range fns {
				if err := fn(slug); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[3].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = func() func(string) error {
		validators := articleDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescDescription is the schema descriptor for description field.
	articleDescDescription := articleFields[4].Descriptor()
	// article.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	article.DescriptionValidator = func() func(string) error {
		validators := articleDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescBody is the schema descriptor for body field.
	articleDescBody := articleFields[5].Descriptor()
	// article.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	article.BodyValidator = func() func(string) error {
		validators := articleDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[6].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[7].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	// articleDescID is the schema descriptor for id field.
	articleDescID := articleFields[0].Descriptor()
	// article.DefaultID holds the default value on creation for the id field.
	article.DefaultID = articleDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescImage is the schema descriptor for image field.
	userDescImage := userFields[4].Descriptor()
	// user.DefaultImage holds the default value on creation for the image field.
	user.DefaultImage = userDescImage.Default.(string)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[5].Descriptor()
	// user.DefaultBio holds the default value on creation for the bio field.
	user.DefaultBio = userDescBio.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	userfollowFields := schema.UserFollow{}.Fields()
	_ = userfollowFields
	// userfollowDescCreatedAt is the schema descriptor for created_at field.
	userfollowDescCreatedAt := userfollowFields[3].Descriptor()
	// userfollow.DefaultCreatedAt holds the default value on creation for the created_at field.
	userfollow.DefaultCreatedAt = userfollowDescCreatedAt.Default.(func() time.Time)
	// userfollowDescID is the schema descriptor for id field.
	userfollowDescID := userfollowFields[0].Descriptor()
	// userfollow.DefaultID holds the default value on creation for the id field.
	userfollow.DefaultID = userfollowDescID.Default.(func() uuid.UUID)
}
