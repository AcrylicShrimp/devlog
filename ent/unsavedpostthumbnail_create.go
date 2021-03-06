// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostthumbnail"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostThumbnailCreate is the builder for creating a UnsavedPostThumbnail entity.
type UnsavedPostThumbnailCreate struct {
	config
	mutation *UnsavedPostThumbnailMutation
	hooks    []Hook
}

// SetValidity sets the "validity" field.
func (uptc *UnsavedPostThumbnailCreate) SetValidity(u unsavedpostthumbnail.Validity) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetValidity(u)
	return uptc
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableValidity(u *unsavedpostthumbnail.Validity) *UnsavedPostThumbnailCreate {
	if u != nil {
		uptc.SetValidity(*u)
	}
	return uptc
}

// SetWidth sets the "width" field.
func (uptc *UnsavedPostThumbnailCreate) SetWidth(u uint32) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetWidth(u)
	return uptc
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableWidth(u *uint32) *UnsavedPostThumbnailCreate {
	if u != nil {
		uptc.SetWidth(*u)
	}
	return uptc
}

// SetHeight sets the "height" field.
func (uptc *UnsavedPostThumbnailCreate) SetHeight(u uint32) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetHeight(u)
	return uptc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableHeight(u *uint32) *UnsavedPostThumbnailCreate {
	if u != nil {
		uptc.SetHeight(*u)
	}
	return uptc
}

// SetHash sets the "hash" field.
func (uptc *UnsavedPostThumbnailCreate) SetHash(s string) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetHash(s)
	return uptc
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableHash(s *string) *UnsavedPostThumbnailCreate {
	if s != nil {
		uptc.SetHash(*s)
	}
	return uptc
}

// SetURL sets the "url" field.
func (uptc *UnsavedPostThumbnailCreate) SetURL(s string) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetURL(s)
	return uptc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableURL(s *string) *UnsavedPostThumbnailCreate {
	if s != nil {
		uptc.SetURL(*s)
	}
	return uptc
}

// SetCreatedAt sets the "created_at" field.
func (uptc *UnsavedPostThumbnailCreate) SetCreatedAt(t time.Time) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetCreatedAt(t)
	return uptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uptc *UnsavedPostThumbnailCreate) SetNillableCreatedAt(t *time.Time) *UnsavedPostThumbnailCreate {
	if t != nil {
		uptc.SetCreatedAt(*t)
	}
	return uptc
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (uptc *UnsavedPostThumbnailCreate) SetUnsavedPostID(id int) *UnsavedPostThumbnailCreate {
	uptc.mutation.SetUnsavedPostID(id)
	return uptc
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (uptc *UnsavedPostThumbnailCreate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostThumbnailCreate {
	return uptc.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostThumbnailMutation object of the builder.
func (uptc *UnsavedPostThumbnailCreate) Mutation() *UnsavedPostThumbnailMutation {
	return uptc.mutation
}

// Save creates the UnsavedPostThumbnail in the database.
func (uptc *UnsavedPostThumbnailCreate) Save(ctx context.Context) (*UnsavedPostThumbnail, error) {
	var (
		err  error
		node *UnsavedPostThumbnail
	)
	uptc.defaults()
	if len(uptc.hooks) == 0 {
		if err = uptc.check(); err != nil {
			return nil, err
		}
		node, err = uptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uptc.check(); err != nil {
				return nil, err
			}
			uptc.mutation = mutation
			node, err = uptc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uptc.hooks) - 1; i >= 0; i-- {
			mut = uptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uptc *UnsavedPostThumbnailCreate) SaveX(ctx context.Context) *UnsavedPostThumbnail {
	v, err := uptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uptc *UnsavedPostThumbnailCreate) defaults() {
	if _, ok := uptc.mutation.Validity(); !ok {
		v := unsavedpostthumbnail.DefaultValidity
		uptc.mutation.SetValidity(v)
	}
	if _, ok := uptc.mutation.CreatedAt(); !ok {
		v := unsavedpostthumbnail.DefaultCreatedAt()
		uptc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uptc *UnsavedPostThumbnailCreate) check() error {
	if _, ok := uptc.mutation.Validity(); !ok {
		return &ValidationError{Name: "validity", err: errors.New("ent: missing required field \"validity\"")}
	}
	if v, ok := uptc.mutation.Validity(); ok {
		if err := unsavedpostthumbnail.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := uptc.mutation.Hash(); ok {
		if err := unsavedpostthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := uptc.mutation.URL(); ok {
		if err := unsavedpostthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := uptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := uptc.mutation.UnsavedPostID(); !ok {
		return &ValidationError{Name: "unsaved_post", err: errors.New("ent: missing required edge \"unsaved_post\"")}
	}
	return nil
}

func (uptc *UnsavedPostThumbnailCreate) sqlSave(ctx context.Context) (*UnsavedPostThumbnail, error) {
	_node, _spec := uptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uptc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uptc *UnsavedPostThumbnailCreate) createSpec() (*UnsavedPostThumbnail, *sqlgraph.CreateSpec) {
	var (
		_node = &UnsavedPostThumbnail{config: uptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: unsavedpostthumbnail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostthumbnail.FieldID,
			},
		}
	)
	if value, ok := uptc.mutation.Validity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostthumbnail.FieldValidity,
		})
		_node.Validity = value
	}
	if value, ok := uptc.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldWidth,
		})
		_node.Width = &value
	}
	if value, ok := uptc.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHeight,
		})
		_node.Height = &value
	}
	if value, ok := uptc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHash,
		})
		_node.Hash = &value
	}
	if value, ok := uptc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldURL,
		})
		_node.URL = &value
	}
	if value, ok := uptc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostthumbnail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := uptc.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   unsavedpostthumbnail.UnsavedPostTable,
			Columns: []string{unsavedpostthumbnail.UnsavedPostColumn},
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
		_node.unsaved_post_thumbnail = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UnsavedPostThumbnailCreateBulk is the builder for creating many UnsavedPostThumbnail entities in bulk.
type UnsavedPostThumbnailCreateBulk struct {
	config
	builders []*UnsavedPostThumbnailCreate
}

// Save creates the UnsavedPostThumbnail entities in the database.
func (uptcb *UnsavedPostThumbnailCreateBulk) Save(ctx context.Context) ([]*UnsavedPostThumbnail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uptcb.builders))
	nodes := make([]*UnsavedPostThumbnail, len(uptcb.builders))
	mutators := make([]Mutator, len(uptcb.builders))
	for i := range uptcb.builders {
		func(i int, root context.Context) {
			builder := uptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UnsavedPostThumbnailMutation)
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
					_, err = mutators[i+1].Mutate(root, uptcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uptcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uptcb *UnsavedPostThumbnailCreateBulk) SaveX(ctx context.Context) []*UnsavedPostThumbnail {
	v, err := uptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
