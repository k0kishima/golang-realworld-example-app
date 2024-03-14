// Code generated by ent, DO NOT EDIT.

package tag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeArticle holds the string denoting the article edge name in mutations.
	EdgeArticle = "article"
	// EdgeTagArticle holds the string denoting the tag_article edge name in mutations.
	EdgeTagArticle = "tag_article"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ArticleTable is the table that holds the article relation/edge. The primary key declared below.
	ArticleTable = "article_tags"
	// ArticleInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticleInverseTable = "articles"
	// TagArticleTable is the table that holds the tag_article relation/edge.
	TagArticleTable = "article_tags"
	// TagArticleInverseTable is the table name for the ArticleTag entity.
	// It exists in this package in order to avoid circular dependency with the "articletag" package.
	TagArticleInverseTable = "article_tags"
	// TagArticleColumn is the table column denoting the tag_article relation/edge.
	TagArticleColumn = "tag_id"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldCreatedAt,
}

var (
	// ArticlePrimaryKey and ArticleColumn2 are the table columns denoting the
	// primary key for the article relation (M2M).
	ArticlePrimaryKey = []string{"article_id", "tag_id"}
)

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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByArticleCount orders the results by article count.
func ByArticleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArticleStep(), opts...)
	}
}

// ByArticle orders the results by article terms.
func ByArticle(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArticleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagArticleCount orders the results by tag_article count.
func ByTagArticleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagArticleStep(), opts...)
	}
}

// ByTagArticle orders the results by tag_article terms.
func ByTagArticle(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagArticleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newArticleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArticleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ArticleTable, ArticlePrimaryKey...),
	)
}
func newTagArticleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagArticleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, TagArticleTable, TagArticleColumn),
	)
}