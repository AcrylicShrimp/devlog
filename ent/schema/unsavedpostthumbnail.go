package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UnsavedPostThumbnail holds the schema definition for the UnsavedPostThumbnail entity.
type UnsavedPostThumbnail struct {
	ent.Schema
}

// Fields of the UnsavedPostThumbnail.
func (UnsavedPostThumbnail) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("validity").Values("pending", "valid", "invalid").Default("pending"),
		field.Uint32("width").Optional().Nillable(),
		field.Uint32("height").Optional().Nillable(),
		field.String("hash").MaxLen(64).Optional().Nillable(),
		field.String("url").MaxLen(512).Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UnsavedPostThumbnail.
func (UnsavedPostThumbnail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unsaved_post", UnsavedPost.Type).Ref("thumbnail").Unique().Required(),
	}
}
