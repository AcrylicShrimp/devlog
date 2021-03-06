// Code generated by entc, DO NOT EDIT.

package adminsession

import (
	"time"
)

const (
	// Label holds the string label denoting the adminsession type in the database.
	Label = "admin_session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the adminsession in the database.
	Table = "admin_sessions"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "admin_sessions"
	// UserInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	UserInverseTable = "admins"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "admin_sessions"
)

// Columns holds all SQL columns for adminsession fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldExpiresAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "admin_sessions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"admin_sessions",
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
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
