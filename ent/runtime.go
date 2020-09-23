// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/mattn/entgo-bbs/ent/entry"
	"github.com/mattn/entgo-bbs/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	entryFields := schema.Entry{}.Fields()
	_ = entryFields
	// entryDescContent is the schema descriptor for content field.
	entryDescContent := entryFields[0].Descriptor()
	// entry.DefaultContent holds the default value on creation for the content field.
	entry.DefaultContent = entryDescContent.Default.(string)
	// entryDescCreatedAt is the schema descriptor for created_at field.
	entryDescCreatedAt := entryFields[1].Descriptor()
	// entry.DefaultCreatedAt holds the default value on creation for the created_at field.
	entry.DefaultCreatedAt = entryDescCreatedAt.Default.(func() time.Time)
}