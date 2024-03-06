package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("username").Unique().NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("image").Default("https://avatars.githubusercontent.com/u/32737308?v=4"),
		field.String("bio").Default(""),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").Unique(),
	}
}
