package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("username").Unique().NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("image").Default("https://api.realworld.io/images/smiley-cyrus.jpeg"),
		field.String("bio").Default(""),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type).StorageKey(edge.Column("author_id")).Immutable().Unique(),
		edge.To("comments", Comment.Type).StorageKey(edge.Column("author_id")).Immutable().Unique(),
		edge.To("favoriteArticles", Article.Type).StorageKey(edge.Table("user_favorites")),
		edge.To("following", User.Type).StorageKey(edge.Table("user_follows"), edge.Columns("follower_id", "followee_id")),
	}
}
