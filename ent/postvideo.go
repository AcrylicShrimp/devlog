// Code generated by entc, DO NOT EDIT.

package ent

import (
	"devlog/ent/post"
	"devlog/ent/postvideo"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// PostVideo is the model entity for the PostVideo schema.
type PostVideo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostVideoQuery when eager-loading is set.
	Edges       PostVideoEdges `json:"edges"`
	post_videos *int
}

// PostVideoEdges holds the relations/edges for other nodes in the graph.
type PostVideoEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostVideoEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Post == nil {
			// The edge post was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PostVideo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case postvideo.FieldID:
			values[i] = new(sql.NullInt64)
		case postvideo.FieldUUID, postvideo.FieldTitle, postvideo.FieldURL:
			values[i] = new(sql.NullString)
		case postvideo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case postvideo.ForeignKeys[0]: // post_videos
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PostVideo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PostVideo fields.
func (pv *PostVideo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case postvideo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = int(value.Int64)
		case postvideo.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				pv.UUID = value.String
			}
		case postvideo.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pv.Title = value.String
			}
		case postvideo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pv.URL = value.String
			}
		case postvideo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pv.CreatedAt = value.Time
			}
		case postvideo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field post_videos", value)
			} else if value.Valid {
				pv.post_videos = new(int)
				*pv.post_videos = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPost queries the "post" edge of the PostVideo entity.
func (pv *PostVideo) QueryPost() *PostQuery {
	return (&PostVideoClient{config: pv.config}).QueryPost(pv)
}

// Update returns a builder for updating this PostVideo.
// Note that you need to call PostVideo.Unwrap() before calling this method if this PostVideo
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PostVideo) Update() *PostVideoUpdateOne {
	return (&PostVideoClient{config: pv.config}).UpdateOne(pv)
}

// Unwrap unwraps the PostVideo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PostVideo) Unwrap() *PostVideo {
	tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PostVideo is not a transactional entity")
	}
	pv.config.driver = tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PostVideo) String() string {
	var builder strings.Builder
	builder.WriteString("PostVideo(")
	builder.WriteString(fmt.Sprintf("id=%v", pv.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(pv.UUID)
	builder.WriteString(", title=")
	builder.WriteString(pv.Title)
	builder.WriteString(", url=")
	builder.WriteString(pv.URL)
	builder.WriteString(", created_at=")
	builder.WriteString(pv.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PostVideos is a parsable slice of PostVideo.
type PostVideos []*PostVideo

func (pv PostVideos) config(cfg config) {
	for _i := range pv {
		pv[_i].config = cfg
	}
}
