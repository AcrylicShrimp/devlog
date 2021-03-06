// Code generated by entc, DO NOT EDIT.

package unsavedpost

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the unsavedpost type in the database.
	Label = "unsaved_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldAccessLevel holds the string denoting the access_level field in the database.
	FieldAccessLevel = "access_level"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeThumbnail holds the string denoting the thumbnail edge name in mutations.
	EdgeThumbnail = "thumbnail"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the unsavedpost in the database.
	Table = "unsaved_posts"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "unsaved_posts"
	// AuthorInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AuthorInverseTable = "admins"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "admin_unsaved_posts"
	// CategoryTable is the table the holds the category relation/edge.
	CategoryTable = "unsaved_posts"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_unsaved_posts"
	// ThumbnailTable is the table the holds the thumbnail relation/edge.
	ThumbnailTable = "unsaved_post_thumbnails"
	// ThumbnailInverseTable is the table name for the UnsavedPostThumbnail entity.
	// It exists in this package in order to avoid circular dependency with the "unsavedpostthumbnail" package.
	ThumbnailInverseTable = "unsaved_post_thumbnails"
	// ThumbnailColumn is the table column denoting the thumbnail relation/edge.
	ThumbnailColumn = "unsaved_post_thumbnail"
	// ImagesTable is the table the holds the images relation/edge.
	ImagesTable = "unsaved_post_images"
	// ImagesInverseTable is the table name for the UnsavedPostImage entity.
	// It exists in this package in order to avoid circular dependency with the "unsavedpostimage" package.
	ImagesInverseTable = "unsaved_post_images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "unsaved_post_images"
	// VideosTable is the table the holds the videos relation/edge.
	VideosTable = "unsaved_post_videos"
	// VideosInverseTable is the table name for the UnsavedPostVideo entity.
	// It exists in this package in order to avoid circular dependency with the "unsavedpostvideo" package.
	VideosInverseTable = "unsaved_post_videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "unsaved_post_videos"
	// AttachmentsTable is the table the holds the attachments relation/edge.
	AttachmentsTable = "unsaved_post_attachments"
	// AttachmentsInverseTable is the table name for the UnsavedPostAttachment entity.
	// It exists in this package in order to avoid circular dependency with the "unsavedpostattachment" package.
	AttachmentsInverseTable = "unsaved_post_attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "unsaved_post_attachments"
)

// Columns holds all SQL columns for unsavedpost fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldSlug,
	FieldAccessLevel,
	FieldTitle,
	FieldContent,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "unsaved_posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"admin_unsaved_posts",
	"category_unsaved_posts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	UUIDValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
)

// AccessLevel defines the type for the "access_level" enum field.
type AccessLevel string

// AccessLevel values.
const (
	AccessLevelPublic   AccessLevel = "public"
	AccessLevelUnlisted AccessLevel = "unlisted"
	AccessLevelPrivate  AccessLevel = "private"
)

func (al AccessLevel) String() string {
	return string(al)
}

// AccessLevelValidator is a validator for the "access_level" field enum values. It is called by the builders before save.
func AccessLevelValidator(al AccessLevel) error {
	switch al {
	case AccessLevelPublic, AccessLevelUnlisted, AccessLevelPrivate:
		return nil
	default:
		return fmt.Errorf("unsavedpost: invalid enum value for access_level field: %q", al)
	}
}
