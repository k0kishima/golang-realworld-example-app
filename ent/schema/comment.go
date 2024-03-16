package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("article_id", uuid.UUID{}).Immutable(),
		field.UUID("author_id", uuid.UUID{}).Immutable(),
		field.String("body").NotEmpty().MaxLen(4096),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("comments").
			Immutable(),
	}
}
