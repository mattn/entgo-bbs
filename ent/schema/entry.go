package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").
			Default(""),
		field.Time("created_at").
			Default(func() time.Time {
				return time.Now()
			}),
	}

}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return nil
}
