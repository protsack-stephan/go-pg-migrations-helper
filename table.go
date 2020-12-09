package pgmigrations

import (
	"fmt"
)

// Table creating sql for db table
type Table struct {
	Name        string
	PrimaryKey  []string
	Columns     []Column
	Indexes     []Index
	ForeignKeys []ForeignKey
	Partition   *Partition
}

// Create generate sql for table creation
func (t *Table) Create() string {
	return fmt.Sprintf("")
}

// Drop generate sql for deleting the table
func (t *Table) Drop() string {
	return fmt.Sprintf("")
}
