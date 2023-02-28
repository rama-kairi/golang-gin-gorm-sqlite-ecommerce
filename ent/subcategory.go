// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent/category"
	"github.com/rama-kairi/blog-api-golang-gin/ent/subcategory"
)

// SubCategory is the model entity for the SubCategory schema.
type SubCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID uuid.UUID `json:"category_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubCategoryQuery when eager-loading is set.
	Edges SubCategoryEdges `json:"edges"`
}

// SubCategoryEdges holds the relations/edges for other nodes in the graph.
type SubCategoryEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubCategoryEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e SubCategoryEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldName:
			values[i] = new(sql.NullString)
		case subcategory.FieldCreatedAt, subcategory.FieldUpdatedAt, subcategory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case subcategory.FieldID, subcategory.FieldCategoryID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SubCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubCategory fields.
func (sc *SubCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sc.ID = *value
			}
		case subcategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case subcategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case subcategory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sc.DeletedAt = new(time.Time)
				*sc.DeletedAt = value.Time
			}
		case subcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sc.Name = value.String
			}
		case subcategory.FieldCategoryID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value != nil {
				sc.CategoryID = *value
			}
		}
	}
	return nil
}

// QueryCategory queries the "category" edge of the SubCategory entity.
func (sc *SubCategory) QueryCategory() *CategoryQuery {
	return NewSubCategoryClient(sc.config).QueryCategory(sc)
}

// QueryProducts queries the "products" edge of the SubCategory entity.
func (sc *SubCategory) QueryProducts() *ProductQuery {
	return NewSubCategoryClient(sc.config).QueryProducts(sc)
}

// Update returns a builder for updating this SubCategory.
// Note that you need to call SubCategory.Unwrap() before calling this method if this SubCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SubCategory) Update() *SubCategoryUpdateOne {
	return NewSubCategoryClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the SubCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SubCategory) Unwrap() *SubCategory {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubCategory is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SubCategory) String() string {
	var builder strings.Builder
	builder.WriteString("SubCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sc.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sc.Name)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.CategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// SubCategories is a parsable slice of SubCategory.
type SubCategories []*SubCategory

func (sc SubCategories) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}