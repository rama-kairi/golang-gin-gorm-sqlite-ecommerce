// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent/category"
	"github.com/rama-kairi/blog-api-golang-gin/ent/predicate"
	"github.com/rama-kairi/blog-api-golang-gin/ent/product"
	"github.com/rama-kairi/blog-api-golang-gin/ent/subcategory"
	"github.com/rama-kairi/blog-api-golang-gin/ent/user"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *ProductUpdate) SetDeletedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableDeletedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *ProductUpdate) ClearDeletedAt() *ProductUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetPrice sets the "price" field.
func (pu *ProductUpdate) SetPrice(i int) *ProductUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(i)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePrice(i *int) *ProductUpdate {
	if i != nil {
		pu.SetPrice(*i)
	}
	return pu
}

// AddPrice adds i to the "price" field.
func (pu *ProductUpdate) AddPrice(i int) *ProductUpdate {
	pu.mutation.AddPrice(i)
	return pu
}

// ClearPrice clears the value of the "price" field.
func (pu *ProductUpdate) ClearPrice() *ProductUpdate {
	pu.mutation.ClearPrice()
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *ProductUpdate) SetUserID(u uuid.UUID) *ProductUpdate {
	pu.mutation.SetUserID(u)
	return pu
}

// SetCategoryID sets the "category_id" field.
func (pu *ProductUpdate) SetCategoryID(u uuid.UUID) *ProductUpdate {
	pu.mutation.SetCategoryID(u)
	return pu
}

// SetSubCategoryID sets the "sub_category_id" field.
func (pu *ProductUpdate) SetSubCategoryID(u uuid.UUID) *ProductUpdate {
	pu.mutation.SetSubCategoryID(u)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *ProductUpdate) SetUser(u *User) *ProductUpdate {
	return pu.SetUserID(u.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (pu *ProductUpdate) SetCategory(c *Category) *ProductUpdate {
	return pu.SetCategoryID(c.ID)
}

// SetSubCategory sets the "sub_category" edge to the SubCategory entity.
func (pu *ProductUpdate) SetSubCategory(s *SubCategory) *ProductUpdate {
	return pu.SetSubCategoryID(s.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *ProductUpdate) ClearUser() *ProductUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearCategory clears the "category" edge to the Category entity.
func (pu *ProductUpdate) ClearCategory() *ProductUpdate {
	pu.mutation.ClearCategory()
	return pu
}

// ClearSubCategory clears the "sub_category" edge to the SubCategory entity.
func (pu *ProductUpdate) ClearSubCategory() *ProductUpdate {
	pu.mutation.ClearSubCategory()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks[int, ProductMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProductUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProductUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.user"`)
	}
	if _, ok := pu.mutation.CategoryID(); pu.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.category"`)
	}
	if _, ok := pu.mutation.SubCategoryID(); pu.mutation.SubCategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.sub_category"`)
	}
	return nil
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(product.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(product.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeInt, value)
	}
	if pu.mutation.PriceCleared() {
		_spec.ClearField(product.FieldPrice, field.TypeInt)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.UserTable,
			Columns: []string{product.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.UserTable,
			Columns: []string{product.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.CategoryTable,
			Columns: []string{product.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.CategoryTable,
			Columns: []string{product.CategoryColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SubCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.SubCategoryTable,
			Columns: []string{product.SubCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SubCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.SubCategoryTable,
			Columns: []string{product.SubCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *ProductUpdateOne) SetDeletedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableDeletedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *ProductUpdateOne) ClearDeletedAt() *ProductUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetPrice sets the "price" field.
func (puo *ProductUpdateOne) SetPrice(i int) *ProductUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(i)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePrice(i *int) *ProductUpdateOne {
	if i != nil {
		puo.SetPrice(*i)
	}
	return puo
}

// AddPrice adds i to the "price" field.
func (puo *ProductUpdateOne) AddPrice(i int) *ProductUpdateOne {
	puo.mutation.AddPrice(i)
	return puo
}

// ClearPrice clears the value of the "price" field.
func (puo *ProductUpdateOne) ClearPrice() *ProductUpdateOne {
	puo.mutation.ClearPrice()
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *ProductUpdateOne) SetUserID(u uuid.UUID) *ProductUpdateOne {
	puo.mutation.SetUserID(u)
	return puo
}

// SetCategoryID sets the "category_id" field.
func (puo *ProductUpdateOne) SetCategoryID(u uuid.UUID) *ProductUpdateOne {
	puo.mutation.SetCategoryID(u)
	return puo
}

// SetSubCategoryID sets the "sub_category_id" field.
func (puo *ProductUpdateOne) SetSubCategoryID(u uuid.UUID) *ProductUpdateOne {
	puo.mutation.SetSubCategoryID(u)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *ProductUpdateOne) SetUser(u *User) *ProductUpdateOne {
	return puo.SetUserID(u.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (puo *ProductUpdateOne) SetCategory(c *Category) *ProductUpdateOne {
	return puo.SetCategoryID(c.ID)
}

// SetSubCategory sets the "sub_category" edge to the SubCategory entity.
func (puo *ProductUpdateOne) SetSubCategory(s *SubCategory) *ProductUpdateOne {
	return puo.SetSubCategoryID(s.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *ProductUpdateOne) ClearUser() *ProductUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearCategory clears the "category" edge to the Category entity.
func (puo *ProductUpdateOne) ClearCategory() *ProductUpdateOne {
	puo.mutation.ClearCategory()
	return puo
}

// ClearSubCategory clears the "sub_category" edge to the SubCategory entity.
func (puo *ProductUpdateOne) ClearSubCategory() *ProductUpdateOne {
	puo.mutation.ClearSubCategory()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	puo.defaults()
	return withHooks[*Product, ProductMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProductUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProductUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.user"`)
	}
	if _, ok := puo.mutation.CategoryID(); puo.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.category"`)
	}
	if _, ok := puo.mutation.SubCategoryID(); puo.mutation.SubCategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.sub_category"`)
	}
	return nil
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(product.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(product.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeInt, value)
	}
	if puo.mutation.PriceCleared() {
		_spec.ClearField(product.FieldPrice, field.TypeInt)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.UserTable,
			Columns: []string{product.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.UserTable,
			Columns: []string{product.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.CategoryTable,
			Columns: []string{product.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.CategoryTable,
			Columns: []string{product.CategoryColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SubCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.SubCategoryTable,
			Columns: []string{product.SubCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SubCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.SubCategoryTable,
			Columns: []string{product.SubCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
