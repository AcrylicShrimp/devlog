// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postattachment"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostAttachmentCreate is the builder for creating a PostAttachment entity.
type PostAttachmentCreate struct {
	config
	mutation *PostAttachmentMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (pac *PostAttachmentCreate) SetUUID(s string) *PostAttachmentCreate {
	pac.mutation.SetUUID(s)
	return pac
}

// SetSize sets the "size" field.
func (pac *PostAttachmentCreate) SetSize(u uint64) *PostAttachmentCreate {
	pac.mutation.SetSize(u)
	return pac
}

// SetName sets the "name" field.
func (pac *PostAttachmentCreate) SetName(s string) *PostAttachmentCreate {
	pac.mutation.SetName(s)
	return pac
}

// SetMime sets the "mime" field.
func (pac *PostAttachmentCreate) SetMime(s string) *PostAttachmentCreate {
	pac.mutation.SetMime(s)
	return pac
}

// SetURL sets the "url" field.
func (pac *PostAttachmentCreate) SetURL(s string) *PostAttachmentCreate {
	pac.mutation.SetURL(s)
	return pac
}

// SetCreatedAt sets the "created_at" field.
func (pac *PostAttachmentCreate) SetCreatedAt(t time.Time) *PostAttachmentCreate {
	pac.mutation.SetCreatedAt(t)
	return pac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pac *PostAttachmentCreate) SetNillableCreatedAt(t *time.Time) *PostAttachmentCreate {
	if t != nil {
		pac.SetCreatedAt(*t)
	}
	return pac
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pac *PostAttachmentCreate) SetPostID(id int) *PostAttachmentCreate {
	pac.mutation.SetPostID(id)
	return pac
}

// SetPost sets the "post" edge to the Post entity.
func (pac *PostAttachmentCreate) SetPost(p *Post) *PostAttachmentCreate {
	return pac.SetPostID(p.ID)
}

// Mutation returns the PostAttachmentMutation object of the builder.
func (pac *PostAttachmentCreate) Mutation() *PostAttachmentMutation {
	return pac.mutation
}

// Save creates the PostAttachment in the database.
func (pac *PostAttachmentCreate) Save(ctx context.Context) (*PostAttachment, error) {
	var (
		err  error
		node *PostAttachment
	)
	pac.defaults()
	if len(pac.hooks) == 0 {
		if err = pac.check(); err != nil {
			return nil, err
		}
		node, err = pac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pac.check(); err != nil {
				return nil, err
			}
			pac.mutation = mutation
			node, err = pac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pac.hooks) - 1; i >= 0; i-- {
			mut = pac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pac *PostAttachmentCreate) SaveX(ctx context.Context) *PostAttachment {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pac *PostAttachmentCreate) defaults() {
	if _, ok := pac.mutation.CreatedAt(); !ok {
		v := postattachment.DefaultCreatedAt()
		pac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *PostAttachmentCreate) check() error {
	if _, ok := pac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if v, ok := pac.mutation.UUID(); ok {
		if err := postattachment.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if _, ok := pac.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New("ent: missing required field \"size\"")}
	}
	if _, ok := pac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pac.mutation.Name(); ok {
		if err := postattachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pac.mutation.Mime(); !ok {
		return &ValidationError{Name: "mime", err: errors.New("ent: missing required field \"mime\"")}
	}
	if v, ok := pac.mutation.Mime(); ok {
		if err := postattachment.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if _, ok := pac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := pac.mutation.URL(); ok {
		if err := postattachment.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pac.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New("ent: missing required edge \"post\"")}
	}
	return nil
}

func (pac *PostAttachmentCreate) sqlSave(ctx context.Context) (*PostAttachment, error) {
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pac *PostAttachmentCreate) createSpec() (*PostAttachment, *sqlgraph.CreateSpec) {
	var (
		_node = &PostAttachment{config: pac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: postattachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postattachment.FieldID,
			},
		}
	)
	if value, ok := pac.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := pac.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postattachment.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := pac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pac.mutation.Mime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldMime,
		})
		_node.Mime = value
	}
	if value, ok := pac.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := pac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postattachment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := pac.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postattachment.PostTable,
			Columns: []string{postattachment.PostColumn},
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
		_node.post_attachments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostAttachmentCreateBulk is the builder for creating many PostAttachment entities in bulk.
type PostAttachmentCreateBulk struct {
	config
	builders []*PostAttachmentCreate
}

// Save creates the PostAttachment entities in the database.
func (pacb *PostAttachmentCreateBulk) Save(ctx context.Context) ([]*PostAttachment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*PostAttachment, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostAttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *PostAttachmentCreateBulk) SaveX(ctx context.Context) []*PostAttachment {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
