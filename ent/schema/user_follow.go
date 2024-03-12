package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type UserFollow struct {
	ent.Schema
}

func (UserFollow) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id"),
		field.UUID("follower_id", uuid.UUID{}).StorageKey("follower_id"),
		field.UUID("followee_id", uuid.UUID{}).StorageKey("followee_id"),
		field.Time("created_at").Default(time.Now),
	}
}

func (UserFollow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_id", "followee_id").Unique(),
	}
}
