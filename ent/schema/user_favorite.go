package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type UserFavorite struct {
	ent.Schema
}

func (UserFavorite) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("article_id", uuid.UUID{}).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (UserFavorite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "article_id").Unique(),
	}
}

func (UserFavorite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Immutable().
			Field("user_id"),
		edge.To("article", Article.Type).
			Unique().
			Required().
			Immutable().
			Field("article_id"),
	}
}
