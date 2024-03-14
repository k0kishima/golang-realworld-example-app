package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ArticleTag struct {
	ent.Schema
}

func (ArticleTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("article_id", uuid.UUID{}).Immutable(),
		field.UUID("tag_id", uuid.UUID{}).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (ArticleTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "tag_id").Unique(),
	}
}

func (ArticleTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type).
			Unique().
			Required().
			Immutable().
			Field("article_id"),
		edge.To("tag", Tag.Type).
			Unique().
			Required().
			Immutable().
			Field("tag_id"),
	}
}