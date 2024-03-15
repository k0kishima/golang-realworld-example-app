// Code generated by ent, DO NOT EDIT.

package userfavorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUserID, v))
}

// ArticleID applies equality check predicate on the "article_id" field. It's identical to ArticleIDEQ.
func ArticleID(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldArticleID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldUserID, vs...))
}

// ArticleIDEQ applies the EQ predicate on the "article_id" field.
func ArticleIDEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldArticleID, v))
}

// ArticleIDNEQ applies the NEQ predicate on the "article_id" field.
func ArticleIDNEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldArticleID, v))
}

// ArticleIDIn applies the In predicate on the "article_id" field.
func ArticleIDIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldArticleID, vs...))
}

// ArticleIDNotIn applies the NotIn predicate on the "article_id" field.
func ArticleIDNotIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldArticleID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArticle applies the HasEdge predicate on the "article" edge.
func HasArticle() predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleWith applies the HasEdge predicate on the "article" edge with a given conditions (other predicates).
func HasArticleWith(preds ...predicate.Article) predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := newArticleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.NotPredicates(p))
}
