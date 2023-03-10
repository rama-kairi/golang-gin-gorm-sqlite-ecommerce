// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt, Nullable: true},
		{Name: "category_id", Type: field.TypeUUID},
		{Name: "sub_category_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_categories_products",
				Columns:    []*schema.Column{ProductsColumns[7]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "products_sub_categories_products",
				Columns:    []*schema.Column{ProductsColumns[8]},
				RefColumns: []*schema.Column{SubCategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "products_users_products",
				Columns:    []*schema.Column{ProductsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubCategoriesColumns holds the columns for the "sub_categories" table.
	SubCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeUUID},
	}
	// SubCategoriesTable holds the schema information for the "sub_categories" table.
	SubCategoriesTable = &schema.Table{
		Name:       "sub_categories",
		Columns:    SubCategoriesColumns,
		PrimaryKey: []*schema.Column{SubCategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sub_categories_categories_sub_categories",
				Columns:    []*schema.Column{SubCategoriesColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		ProductsTable,
		SubCategoriesTable,
		UsersTable,
	}
)

func init() {
	ProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	ProductsTable.ForeignKeys[1].RefTable = SubCategoriesTable
	ProductsTable.ForeignKeys[2].RefTable = UsersTable
	SubCategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
}
