// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostDelete is the builder for deleting a UnsavedPost entity.
type UnsavedPostDelete struct {
	config
	hooks    []Hook
	mutation *UnsavedPostMutation
}

// Where adds a new predicate to the UnsavedPostDelete builder.
func (upd *UnsavedPostDelete) Where(ps ...predicate.UnsavedPost) *UnsavedPostDelete {
	upd.mutation.predicates = append(upd.mutation.predicates, ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UnsavedPostDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upd.hooks) == 0 {
		affected, err = upd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upd.mutation = mutation
			affected, err = upd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upd.hooks) - 1; i >= 0; i-- {
			mut = upd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UnsavedPostDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UnsavedPostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: unsavedpost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpost.FieldID,
			},
		},
	}
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
}

// UnsavedPostDeleteOne is the builder for deleting a single UnsavedPost entity.
type UnsavedPostDeleteOne struct {
	upd *UnsavedPostDelete
}

// Exec executes the deletion query.
func (updo *UnsavedPostDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{unsavedpost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UnsavedPostDeleteOne) ExecX(ctx context.Context) {
	updo.upd.ExecX(ctx)
}
