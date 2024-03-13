package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().StorageKey("id"),
		field.UUID("author_id", uuid.UUID{}).Immutable().StorageKey("author_id"),
		field.String("slug").Unique().NotEmpty().MaxLen(255),
		field.String("title").NotEmpty().MaxLen(255),
		field.String("description").NotEmpty().MaxLen(255),
		field.String("body").NotEmpty().MaxLen(4096),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}
