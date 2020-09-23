package schema

import "github.com/facebook/ent"

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return nil
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return nil
}
