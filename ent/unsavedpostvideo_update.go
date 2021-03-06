// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostvideo"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostVideoUpdate is the builder for updating UnsavedPostVideo entities.
type UnsavedPostVideoUpdate struct {
	config
	hooks    []Hook
	mutation *UnsavedPostVideoMutation
}

// Where adds a new predicate for the UnsavedPostVideoUpdate builder.
func (upvu *UnsavedPostVideoUpdate) Where(ps ...predicate.UnsavedPostVideo) *UnsavedPostVideoUpdate {
	upvu.mutation.predicates = append(upvu.mutation.predicates, ps...)
	return upvu
}

// SetUUID sets the "uuid" field.
func (upvu *UnsavedPostVideoUpdate) SetUUID(s string) *UnsavedPostVideoUpdate {
	upvu.mutation.SetUUID(s)
	return upvu
}

// SetValidity sets the "validity" field.
func (upvu *UnsavedPostVideoUpdate) SetValidity(u unsavedpostvideo.Validity) *UnsavedPostVideoUpdate {
	upvu.mutation.SetValidity(u)
	return upvu
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upvu *UnsavedPostVideoUpdate) SetNillableValidity(u *unsavedpostvideo.Validity) *UnsavedPostVideoUpdate {
	if u != nil {
		upvu.SetValidity(*u)
	}
	return upvu
}

// SetTitle sets the "title" field.
func (upvu *UnsavedPostVideoUpdate) SetTitle(s string) *UnsavedPostVideoUpdate {
	upvu.mutation.SetTitle(s)
	return upvu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (upvu *UnsavedPostVideoUpdate) SetNillableTitle(s *string) *UnsavedPostVideoUpdate {
	if s != nil {
		upvu.SetTitle(*s)
	}
	return upvu
}

// ClearTitle clears the value of the "title" field.
func (upvu *UnsavedPostVideoUpdate) ClearTitle() *UnsavedPostVideoUpdate {
	upvu.mutation.ClearTitle()
	return upvu
}

// SetURL sets the "url" field.
func (upvu *UnsavedPostVideoUpdate) SetURL(s string) *UnsavedPostVideoUpdate {
	upvu.mutation.SetURL(s)
	return upvu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upvu *UnsavedPostVideoUpdate) SetNillableURL(s *string) *UnsavedPostVideoUpdate {
	if s != nil {
		upvu.SetURL(*s)
	}
	return upvu
}

// ClearURL clears the value of the "url" field.
func (upvu *UnsavedPostVideoUpdate) ClearURL() *UnsavedPostVideoUpdate {
	upvu.mutation.ClearURL()
	return upvu
}

// SetCreatedAt sets the "created_at" field.
func (upvu *UnsavedPostVideoUpdate) SetCreatedAt(t time.Time) *UnsavedPostVideoUpdate {
	upvu.mutation.SetCreatedAt(t)
	return upvu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upvu *UnsavedPostVideoUpdate) SetNillableCreatedAt(t *time.Time) *UnsavedPostVideoUpdate {
	if t != nil {
		upvu.SetCreatedAt(*t)
	}
	return upvu
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upvu *UnsavedPostVideoUpdate) SetUnsavedPostID(id int) *UnsavedPostVideoUpdate {
	upvu.mutation.SetUnsavedPostID(id)
	return upvu
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upvu *UnsavedPostVideoUpdate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostVideoUpdate {
	return upvu.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostVideoMutation object of the builder.
func (upvu *UnsavedPostVideoUpdate) Mutation() *UnsavedPostVideoMutation {
	return upvu.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upvu *UnsavedPostVideoUpdate) ClearUnsavedPost() *UnsavedPostVideoUpdate {
	upvu.mutation.ClearUnsavedPost()
	return upvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upvu *UnsavedPostVideoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upvu.hooks) == 0 {
		if err = upvu.check(); err != nil {
			return 0, err
		}
		affected, err = upvu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upvu.check(); err != nil {
				return 0, err
			}
			upvu.mutation = mutation
			affected, err = upvu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upvu.hooks) - 1; i >= 0; i-- {
			mut = upvu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upvu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (upvu *UnsavedPostVideoUpdate) SaveX(ctx context.Context) int {
	affected, err := upvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upvu *UnsavedPostVideoUpdate) Exec(ctx context.Context) error {
	_, err := upvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upvu *UnsavedPostVideoUpdate) ExecX(ctx context.Context) {
	if err := upvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upvu *UnsavedPostVideoUpdate) check() error {
	if v, ok := upvu.mutation.UUID(); ok {
		if err := unsavedpostvideo.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upvu.mutation.Validity(); ok {
		if err := unsavedpostvideo.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upvu.mutation.Title(); ok {
		if err := unsavedpostvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := upvu.mutation.URL(); ok {
		if err := unsavedpostvideo.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upvu.mutation.UnsavedPostID(); upvu.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upvu *UnsavedPostVideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostvideo.Table,
			Columns: unsavedpostvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostvideo.FieldID,
			},
		},
	}
	if ps := upvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upvu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldUUID,
		})
	}
	if value, ok := upvu.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostvideo.FieldValidity,
		})
	}
	if value, ok := upvu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldTitle,
		})
	}
	if upvu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostvideo.FieldTitle,
		})
	}
	if value, ok := upvu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldURL,
		})
	}
	if upvu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostvideo.FieldURL,
		})
	}
	if value, ok := upvu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostvideo.FieldCreatedAt,
		})
	}
	if upvu.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostvideo.UnsavedPostTable,
			Columns: []string{unsavedpostvideo.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upvu.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostvideo.UnsavedPostTable,
			Columns: []string{unsavedpostvideo.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostvideo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UnsavedPostVideoUpdateOne is the builder for updating a single UnsavedPostVideo entity.
type UnsavedPostVideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnsavedPostVideoMutation
}

// SetUUID sets the "uuid" field.
func (upvuo *UnsavedPostVideoUpdateOne) SetUUID(s string) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetUUID(s)
	return upvuo
}

// SetValidity sets the "validity" field.
func (upvuo *UnsavedPostVideoUpdateOne) SetValidity(u unsavedpostvideo.Validity) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetValidity(u)
	return upvuo
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upvuo *UnsavedPostVideoUpdateOne) SetNillableValidity(u *unsavedpostvideo.Validity) *UnsavedPostVideoUpdateOne {
	if u != nil {
		upvuo.SetValidity(*u)
	}
	return upvuo
}

// SetTitle sets the "title" field.
func (upvuo *UnsavedPostVideoUpdateOne) SetTitle(s string) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetTitle(s)
	return upvuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (upvuo *UnsavedPostVideoUpdateOne) SetNillableTitle(s *string) *UnsavedPostVideoUpdateOne {
	if s != nil {
		upvuo.SetTitle(*s)
	}
	return upvuo
}

// ClearTitle clears the value of the "title" field.
func (upvuo *UnsavedPostVideoUpdateOne) ClearTitle() *UnsavedPostVideoUpdateOne {
	upvuo.mutation.ClearTitle()
	return upvuo
}

// SetURL sets the "url" field.
func (upvuo *UnsavedPostVideoUpdateOne) SetURL(s string) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetURL(s)
	return upvuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upvuo *UnsavedPostVideoUpdateOne) SetNillableURL(s *string) *UnsavedPostVideoUpdateOne {
	if s != nil {
		upvuo.SetURL(*s)
	}
	return upvuo
}

// ClearURL clears the value of the "url" field.
func (upvuo *UnsavedPostVideoUpdateOne) ClearURL() *UnsavedPostVideoUpdateOne {
	upvuo.mutation.ClearURL()
	return upvuo
}

// SetCreatedAt sets the "created_at" field.
func (upvuo *UnsavedPostVideoUpdateOne) SetCreatedAt(t time.Time) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetCreatedAt(t)
	return upvuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upvuo *UnsavedPostVideoUpdateOne) SetNillableCreatedAt(t *time.Time) *UnsavedPostVideoUpdateOne {
	if t != nil {
		upvuo.SetCreatedAt(*t)
	}
	return upvuo
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upvuo *UnsavedPostVideoUpdateOne) SetUnsavedPostID(id int) *UnsavedPostVideoUpdateOne {
	upvuo.mutation.SetUnsavedPostID(id)
	return upvuo
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upvuo *UnsavedPostVideoUpdateOne) SetUnsavedPost(u *UnsavedPost) *UnsavedPostVideoUpdateOne {
	return upvuo.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostVideoMutation object of the builder.
func (upvuo *UnsavedPostVideoUpdateOne) Mutation() *UnsavedPostVideoMutation {
	return upvuo.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upvuo *UnsavedPostVideoUpdateOne) ClearUnsavedPost() *UnsavedPostVideoUpdateOne {
	upvuo.mutation.ClearUnsavedPost()
	return upvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upvuo *UnsavedPostVideoUpdateOne) Select(field string, fields ...string) *UnsavedPostVideoUpdateOne {
	upvuo.fields = append([]string{field}, fields...)
	return upvuo
}

// Save executes the query and returns the updated UnsavedPostVideo entity.
func (upvuo *UnsavedPostVideoUpdateOne) Save(ctx context.Context) (*UnsavedPostVideo, error) {
	var (
		err  error
		node *UnsavedPostVideo
	)
	if len(upvuo.hooks) == 0 {
		if err = upvuo.check(); err != nil {
			return nil, err
		}
		node, err = upvuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upvuo.check(); err != nil {
				return nil, err
			}
			upvuo.mutation = mutation
			node, err = upvuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upvuo.hooks) - 1; i >= 0; i-- {
			mut = upvuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upvuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (upvuo *UnsavedPostVideoUpdateOne) SaveX(ctx context.Context) *UnsavedPostVideo {
	node, err := upvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upvuo *UnsavedPostVideoUpdateOne) Exec(ctx context.Context) error {
	_, err := upvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upvuo *UnsavedPostVideoUpdateOne) ExecX(ctx context.Context) {
	if err := upvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upvuo *UnsavedPostVideoUpdateOne) check() error {
	if v, ok := upvuo.mutation.UUID(); ok {
		if err := unsavedpostvideo.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upvuo.mutation.Validity(); ok {
		if err := unsavedpostvideo.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upvuo.mutation.Title(); ok {
		if err := unsavedpostvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := upvuo.mutation.URL(); ok {
		if err := unsavedpostvideo.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upvuo.mutation.UnsavedPostID(); upvuo.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upvuo *UnsavedPostVideoUpdateOne) sqlSave(ctx context.Context) (_node *UnsavedPostVideo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostvideo.Table,
			Columns: unsavedpostvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostvideo.FieldID,
			},
		},
	}
	id, ok := upvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UnsavedPostVideo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := upvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostvideo.FieldID)
		for _, f := range fields {
			if !unsavedpostvideo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unsavedpostvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upvuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldUUID,
		})
	}
	if value, ok := upvuo.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostvideo.FieldValidity,
		})
	}
	if value, ok := upvuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldTitle,
		})
	}
	if upvuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostvideo.FieldTitle,
		})
	}
	if value, ok := upvuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldURL,
		})
	}
	if upvuo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostvideo.FieldURL,
		})
	}
	if value, ok := upvuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostvideo.FieldCreatedAt,
		})
	}
	if upvuo.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostvideo.UnsavedPostTable,
			Columns: []string{unsavedpostvideo.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upvuo.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostvideo.UnsavedPostTable,
			Columns: []string{unsavedpostvideo.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UnsavedPostVideo{config: upvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostvideo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
