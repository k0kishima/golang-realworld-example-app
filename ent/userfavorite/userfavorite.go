// Code generated by ent, DO NOT EDIT.

package userfavorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userfavorite type in the database.
	Label = "user_favorite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldArticleID holds the string denoting the article_id field in the database.
	FieldArticleID = "article_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the userfavorite in the database.
	Table = "user_favorites"
)

// Columns holds all SQL columns for userfavorite fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldArticleID,
	FieldCreatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the UserFavorite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByArticleID orders the results by the article_id field.
func ByArticleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArticleID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}