// Code generated by ent, DO NOT EDIT.

package category

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSubCategories holds the string denoting the sub_categories edge name in mutations.
	EdgeSubCategories = "sub_categories"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// SubCategoriesTable is the table that holds the sub_categories relation/edge.
	SubCategoriesTable = "sub_categories"
	// SubCategoriesInverseTable is the table name for the SubCategory entity.
	// It exists in this package in order to avoid circular dependency with the "subcategory" package.
	SubCategoriesInverseTable = "sub_categories"
	// SubCategoriesColumn is the table column denoting the sub_categories relation/edge.
	SubCategoriesColumn = "category_id"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "category_id"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
