package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("description").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("tags").
			Through("tag_article", ArticleTag.Type),
	}
}
