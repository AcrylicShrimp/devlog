package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64).Unique(),
		field.String("slug").MaxLen(255).Unique(),
		field.Enum("access_level").Values("public", "unlisted", "private"),
		field.String("title").MaxLen(255),
		field.Text("content"),
		field.Text("html_content"),
		field.String("preview_content").MaxLen(255),
		field.Time("created_at").Default(time.Now),
		field.Time("modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Admin.Type).Ref("posts").Unique().Required(),
		edge.From("category", Category.Type).Ref("posts").Unique(),
		edge.To("thumbnail", PostThumbnail.Type).Unique().Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("images", PostImage.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("videos", PostVideo.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("attachments", PostAttachment.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
