// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postthumbnail"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostThumbnailUpdate is the builder for updating PostThumbnail entities.
type PostThumbnailUpdate struct {
	config
	hooks    []Hook
	mutation *PostThumbnailMutation
}

// Where adds a new predicate for the PostThumbnailUpdate builder.
func (ptu *PostThumbnailUpdate) Where(ps ...predicate.PostThumbnail) *PostThumbnailUpdate {
	ptu.mutation.predicates = append(ptu.mutation.predicates, ps...)
	return ptu
}

// SetWidth sets the "width" field.
func (ptu *PostThumbnailUpdate) SetWidth(u uint32) *PostThumbnailUpdate {
	ptu.mutation.ResetWidth()
	ptu.mutation.SetWidth(u)
	return ptu
}

// AddWidth adds u to the "width" field.
func (ptu *PostThumbnailUpdate) AddWidth(u uint32) *PostThumbnailUpdate {
	ptu.mutation.AddWidth(u)
	return ptu
}

// SetHeight sets the "height" field.
func (ptu *PostThumbnailUpdate) SetHeight(u uint32) *PostThumbnailUpdate {
	ptu.mutation.ResetHeight()
	ptu.mutation.SetHeight(u)
	return ptu
}

// AddHeight adds u to the "height" field.
func (ptu *PostThumbnailUpdate) AddHeight(u uint32) *PostThumbnailUpdate {
	ptu.mutation.AddHeight(u)
	return ptu
}

// SetHash sets the "hash" field.
func (ptu *PostThumbnailUpdate) SetHash(s string) *PostThumbnailUpdate {
	ptu.mutation.SetHash(s)
	return ptu
}

// SetURL sets the "url" field.
func (ptu *PostThumbnailUpdate) SetURL(s string) *PostThumbnailUpdate {
	ptu.mutation.SetURL(s)
	return ptu
}

// SetCreatedAt sets the "created_at" field.
func (ptu *PostThumbnailUpdate) SetCreatedAt(t time.Time) *PostThumbnailUpdate {
	ptu.mutation.SetCreatedAt(t)
	return ptu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptu *PostThumbnailUpdate) SetNillableCreatedAt(t *time.Time) *PostThumbnailUpdate {
	if t != nil {
		ptu.SetCreatedAt(*t)
	}
	return ptu
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (ptu *PostThumbnailUpdate) SetPostID(id int) *PostThumbnailUpdate {
	ptu.mutation.SetPostID(id)
	return ptu
}

// SetPost sets the "post" edge to the Post entity.
func (ptu *PostThumbnailUpdate) SetPost(p *Post) *PostThumbnailUpdate {
	return ptu.SetPostID(p.ID)
}

// Mutation returns the PostThumbnailMutation object of the builder.
func (ptu *PostThumbnailUpdate) Mutation() *PostThumbnailMutation {
	return ptu.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (ptu *PostThumbnailUpdate) ClearPost() *PostThumbnailUpdate {
	ptu.mutation.ClearPost()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PostThumbnailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptu.hooks) == 0 {
		if err = ptu.check(); err != nil {
			return 0, err
		}
		affected, err = ptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptu.check(); err != nil {
				return 0, err
			}
			ptu.mutation = mutation
			affected, err = ptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptu.hooks) - 1; i >= 0; i-- {
			mut = ptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PostThumbnailUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PostThumbnailUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PostThumbnailUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *PostThumbnailUpdate) check() error {
	if v, ok := ptu.mutation.Hash(); ok {
		if err := postthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := ptu.mutation.URL(); ok {
		if err := postthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := ptu.mutation.PostID(); ptu.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (ptu *PostThumbnailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postthumbnail.Table,
			Columns: postthumbnail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postthumbnail.FieldID,
			},
		},
	}
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldWidth,
		})
	}
	if value, ok := ptu.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldWidth,
		})
	}
	if value, ok := ptu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldHeight,
		})
	}
	if value, ok := ptu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldHeight,
		})
	}
	if value, ok := ptu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldHash,
		})
	}
	if value, ok := ptu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldURL,
		})
	}
	if value, ok := ptu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postthumbnail.FieldCreatedAt,
		})
	}
	if ptu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   postthumbnail.PostTable,
			Columns: []string{postthumbnail.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   postthumbnail.PostTable,
			Columns: []string{postthumbnail.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postthumbnail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PostThumbnailUpdateOne is the builder for updating a single PostThumbnail entity.
type PostThumbnailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostThumbnailMutation
}

// SetWidth sets the "width" field.
func (ptuo *PostThumbnailUpdateOne) SetWidth(u uint32) *PostThumbnailUpdateOne {
	ptuo.mutation.ResetWidth()
	ptuo.mutation.SetWidth(u)
	return ptuo
}

// AddWidth adds u to the "width" field.
func (ptuo *PostThumbnailUpdateOne) AddWidth(u uint32) *PostThumbnailUpdateOne {
	ptuo.mutation.AddWidth(u)
	return ptuo
}

// SetHeight sets the "height" field.
func (ptuo *PostThumbnailUpdateOne) SetHeight(u uint32) *PostThumbnailUpdateOne {
	ptuo.mutation.ResetHeight()
	ptuo.mutation.SetHeight(u)
	return ptuo
}

// AddHeight adds u to the "height" field.
func (ptuo *PostThumbnailUpdateOne) AddHeight(u uint32) *PostThumbnailUpdateOne {
	ptuo.mutation.AddHeight(u)
	return ptuo
}

// SetHash sets the "hash" field.
func (ptuo *PostThumbnailUpdateOne) SetHash(s string) *PostThumbnailUpdateOne {
	ptuo.mutation.SetHash(s)
	return ptuo
}

// SetURL sets the "url" field.
func (ptuo *PostThumbnailUpdateOne) SetURL(s string) *PostThumbnailUpdateOne {
	ptuo.mutation.SetURL(s)
	return ptuo
}

// SetCreatedAt sets the "created_at" field.
func (ptuo *PostThumbnailUpdateOne) SetCreatedAt(t time.Time) *PostThumbnailUpdateOne {
	ptuo.mutation.SetCreatedAt(t)
	return ptuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptuo *PostThumbnailUpdateOne) SetNillableCreatedAt(t *time.Time) *PostThumbnailUpdateOne {
	if t != nil {
		ptuo.SetCreatedAt(*t)
	}
	return ptuo
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (ptuo *PostThumbnailUpdateOne) SetPostID(id int) *PostThumbnailUpdateOne {
	ptuo.mutation.SetPostID(id)
	return ptuo
}

// SetPost sets the "post" edge to the Post entity.
func (ptuo *PostThumbnailUpdateOne) SetPost(p *Post) *PostThumbnailUpdateOne {
	return ptuo.SetPostID(p.ID)
}

// Mutation returns the PostThumbnailMutation object of the builder.
func (ptuo *PostThumbnailUpdateOne) Mutation() *PostThumbnailMutation {
	return ptuo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (ptuo *PostThumbnailUpdateOne) ClearPost() *PostThumbnailUpdateOne {
	ptuo.mutation.ClearPost()
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PostThumbnailUpdateOne) Select(field string, fields ...string) *PostThumbnailUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PostThumbnail entity.
func (ptuo *PostThumbnailUpdateOne) Save(ctx context.Context) (*PostThumbnail, error) {
	var (
		err  error
		node *PostThumbnail
	)
	if len(ptuo.hooks) == 0 {
		if err = ptuo.check(); err != nil {
			return nil, err
		}
		node, err = ptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptuo.check(); err != nil {
				return nil, err
			}
			ptuo.mutation = mutation
			node, err = ptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptuo.hooks) - 1; i >= 0; i-- {
			mut = ptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PostThumbnailUpdateOne) SaveX(ctx context.Context) *PostThumbnail {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PostThumbnailUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PostThumbnailUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *PostThumbnailUpdateOne) check() error {
	if v, ok := ptuo.mutation.Hash(); ok {
		if err := postthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := ptuo.mutation.URL(); ok {
		if err := postthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := ptuo.mutation.PostID(); ptuo.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (ptuo *PostThumbnailUpdateOne) sqlSave(ctx context.Context) (_node *PostThumbnail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postthumbnail.Table,
			Columns: postthumbnail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postthumbnail.FieldID,
			},
		},
	}
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PostThumbnail.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postthumbnail.FieldID)
		for _, f := range fields {
			if !postthumbnail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != postthumbnail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldWidth,
		})
	}
	if value, ok := ptuo.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldWidth,
		})
	}
	if value, ok := ptuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldHeight,
		})
	}
	if value, ok := ptuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldHeight,
		})
	}
	if value, ok := ptuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldHash,
		})
	}
	if value, ok := ptuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldURL,
		})
	}
	if value, ok := ptuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postthumbnail.FieldCreatedAt,
		})
	}
	if ptuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   postthumbnail.PostTable,
			Columns: []string{postthumbnail.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   postthumbnail.PostTable,
			Columns: []string{postthumbnail.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PostThumbnail{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postthumbnail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
