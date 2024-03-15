// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeFollows holds the string denoting the follows edge name in mutations.
	EdgeFollows = "follows"
	// EdgeArticles holds the string denoting the articles edge name in mutations.
	EdgeArticles = "articles"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeFavorites holds the string denoting the favorites edge name in mutations.
	EdgeFavorites = "favorites"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FollowsTable is the table that holds the follows relation/edge.
	FollowsTable = "user_follows"
	// FollowsInverseTable is the table name for the UserFollow entity.
	// It exists in this package in order to avoid circular dependency with the "userfollow" package.
	FollowsInverseTable = "user_follows"
	// FollowsColumn is the table column denoting the follows relation/edge.
	FollowsColumn = "follower_id"
	// ArticlesTable is the table that holds the articles relation/edge.
	ArticlesTable = "articles"
	// ArticlesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticlesInverseTable = "articles"
	// ArticlesColumn is the table column denoting the articles relation/edge.
	ArticlesColumn = "author_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "author_id"
	// FavoritesTable is the table that holds the favorites relation/edge.
	FavoritesTable = "user_favorites"
	// FavoritesInverseTable is the table name for the UserFavorite entity.
	// It exists in this package in order to avoid circular dependency with the "userfavorite" package.
	FavoritesInverseTable = "user_favorites"
	// FavoritesColumn is the table column denoting the favorites relation/edge.
	FavoritesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldImage,
	FieldBio,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultImage holds the default value on creation for the "image" field.
	DefaultImage string
	// DefaultBio holds the default value on creation for the "bio" field.
	DefaultBio string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByFollowsCount orders the results by follows count.
func ByFollowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowsStep(), opts...)
	}
}

// ByFollows orders the results by follows terms.
func ByFollows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByArticlesCount orders the results by articles count.
func ByArticlesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArticlesStep(), opts...)
	}
}

// ByArticles orders the results by articles terms.
func ByArticles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArticlesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFavoritesCount orders the results by favorites count.
func ByFavoritesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavoritesStep(), opts...)
	}
}

// ByFavorites orders the results by favorites terms.
func ByFavorites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavoritesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFollowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowsTable, FollowsColumn),
	)
}
func newArticlesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArticlesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ArticlesTable, ArticlesColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newFavoritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavoritesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
	)
}
