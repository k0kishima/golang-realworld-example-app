package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (UserFollow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_id", "followee_id").Unique(),
	}
}

func (UserFollow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", User.Type).
			Ref("follows").
			Unique().
			Required().
			Field("follower_id"),
		edge.To("followee", User.Type).
			Unique().
			Required().
			Field("followee_id"),
	}
}
