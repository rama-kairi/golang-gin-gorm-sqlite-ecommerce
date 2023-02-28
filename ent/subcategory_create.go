// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent/category"
	"github.com/rama-kairi/blog-api-golang-gin/ent/product"
	"github.com/rama-kairi/blog-api-golang-gin/ent/subcategory"
)

// SubCategoryCreate is the builder for creating a SubCategory entity.
type SubCategoryCreate struct {
	config
	mutation *SubCategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *SubCategoryCreate) SetCreatedAt(t time.Time) *SubCategoryCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *SubCategoryCreate) SetNillableCreatedAt(t *time.Time) *SubCategoryCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *SubCategoryCreate) SetUpdatedAt(t time.Time) *SubCategoryCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *SubCategoryCreate) SetNillableUpdatedAt(t *time.Time) *SubCategoryCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetDeletedAt sets the "deleted_at" field.
func (scc *SubCategoryCreate) SetDeletedAt(t time.Time) *SubCategoryCreate {
	scc.mutation.SetDeletedAt(t)
	return scc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (scc *SubCategoryCreate) SetNillableDeletedAt(t *time.Time) *SubCategoryCreate {
	if t != nil {
		scc.SetDeletedAt(*t)
	}
	return scc
}

// SetName sets the "name" field.
func (scc *SubCategoryCreate) SetName(s string) *SubCategoryCreate {
	scc.mutation.SetName(s)
	return scc
}

// SetCategoryID sets the "category_id" field.
func (scc *SubCategoryCreate) SetCategoryID(u uuid.UUID) *SubCategoryCreate {
	scc.mutation.SetCategoryID(u)
	return scc
}

// SetID sets the "id" field.
func (scc *SubCategoryCreate) SetID(u uuid.UUID) *SubCategoryCreate {
	scc.mutation.SetID(u)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *SubCategoryCreate) SetNillableID(u *uuid.UUID) *SubCategoryCreate {
	if u != nil {
		scc.SetID(*u)
	}
	return scc
}

// SetCategory sets the "category" edge to the Category entity.
func (scc *SubCategoryCreate) SetCategory(c *Category) *SubCategoryCreate {
	return scc.SetCategoryID(c.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (scc *SubCategoryCreate) AddProductIDs(ids ...uuid.UUID) *SubCategoryCreate {
	scc.mutation.AddProductIDs(ids...)
	return scc
}

// AddProducts adds the "products" edges to the Product entity.
func (scc *SubCategoryCreate) AddProducts(p ...*Product) *SubCategoryCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scc.AddProductIDs(ids...)
}

// Mutation returns the SubCategoryMutation object of the builder.
func (scc *SubCategoryCreate) Mutation() *SubCategoryMutation {
	return scc.mutation
}

// Save creates the SubCategory in the database.
func (scc *SubCategoryCreate) Save(ctx context.Context) (*SubCategory, error) {
	scc.defaults()
	return withHooks[*SubCategory, SubCategoryMutation](ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SubCategoryCreate) SaveX(ctx context.Context) *SubCategory {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SubCategoryCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SubCategoryCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SubCategoryCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := subcategory.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := subcategory.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.ID(); !ok {
		v := subcategory.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SubCategoryCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SubCategory.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SubCategory.updated_at"`)}
	}
	if _, ok := scc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SubCategory.name"`)}
	}
	if v, ok := scc.mutation.Name(); ok {
		if err := subcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SubCategory.name": %w`, err)}
		}
	}
	if _, ok := scc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "SubCategory.category_id"`)}
	}
	if _, ok := scc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "SubCategory.category"`)}
	}
	return nil
}

func (scc *SubCategoryCreate) sqlSave(ctx context.Context) (*SubCategory, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SubCategoryCreate) createSpec() (*SubCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &SubCategory{config: scc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subcategory.FieldID,
			},
		}
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(subcategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.SetField(subcategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.DeletedAt(); ok {
		_spec.SetField(subcategory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := scc.mutation.Name(); ok {
		_spec.SetField(subcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := scc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductsTable,
			Columns: []string{subcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubCategoryCreateBulk is the builder for creating many SubCategory entities in bulk.
type SubCategoryCreateBulk struct {
	config
	builders []*SubCategoryCreate
}

// Save creates the SubCategory entities in the database.
func (sccb *SubCategoryCreateBulk) Save(ctx context.Context) ([]*SubCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SubCategory, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubCategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SubCategoryCreateBulk) SaveX(ctx context.Context) []*SubCategory {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SubCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SubCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
