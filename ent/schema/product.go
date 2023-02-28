package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Int("price").Optional(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("category_id", uuid.UUID{}),
		field.UUID("sub_category_id", uuid.UUID{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("products").Unique().Required().Field("user_id"),
		edge.From("category", Category.Type).Ref("products").Unique().Required().Field("category_id"),
		edge.From("sub_category", SubCategory.Type).Ref("products").Unique().Required().Field("sub_category_id"),
	}
}
