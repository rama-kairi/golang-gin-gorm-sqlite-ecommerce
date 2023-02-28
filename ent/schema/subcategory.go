package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SubCategory holds the schema definition for the SubCategory entity.
type SubCategory struct {
	ent.Schema
}

func (SubCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the SubCategory.
func (SubCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.UUID("category_id", uuid.UUID{}),
	}
}

// Edges of the SubCategory.
func (SubCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("sub_categories").
			Unique().
			Required().
			Field("category_id"),
		edge.To("products", Product.Type),
	}
}
