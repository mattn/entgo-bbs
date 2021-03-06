// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// EntriesColumns holds the columns for the "entries" table.
	EntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
	}
	// EntriesTable holds the schema information for the "entries" table.
	EntriesTable = &schema.Table{
		Name:        "entries",
		Columns:     EntriesColumns,
		PrimaryKey:  []*schema.Column{EntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EntriesTable,
	}
)

func init() {
}
