package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.Int("age").Optional(),
		field.String("address").Optional(),
		field.Bool("is_active").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}
