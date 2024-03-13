package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ArticleTag struct {
	ent.Schema
}

func (ArticleTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("article_id", uuid.UUID{}),
		field.UUID("tag_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (ArticleTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "tag_id").Unique(),
	}
}
