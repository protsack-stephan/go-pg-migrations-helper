package pgmigrations

import (
	"fmt"
	"strings"
)

// Constraint restricting data modification
type Constraint string

// ConstraintNotNull ensures that values in a column cannot be NULL
const ConstraintNotNull Constraint = "NOT NULL"

// ConstraintUnique ensures the values in a column unique across the rows within the same table
const ConstraintUnique Constraint = "UNIQUE"

// ConstraintPrimaryKey uniquely identify rows in a table
const ConstraintPrimaryKey Constraint = "PRIMARY KEY"

// ConstraintCheck constraint ensures the data must satisfy a boolean expression
const ConstraintCheck Constraint = "CHECK"

// ConstraintForeignKey  ensures values in a column or a group of columns from a table
// exists in a column or group of columns in another table
const ConstraintForeignKey Constraint = "FOREIGN KEY"

// Table creating sql for db table
type Table struct {
	Name        string
	Constraints map[Constraint][]string
	Columns     []Column
	Indexes     []Index
	ForeignKeys []ForeignKey
	Partition   *Partition
}

// Create generate sql for table creation
func (t *Table) Create() string {
	columns := ""

	for _, column := range t.Columns {
		columns += fmt.Sprintf("%s,", column.Define())
	}

	for constraint, definitions := range t.Constraints {
		for _, definition := range definitions {
			columns += fmt.Sprintf("%s %s,", constraint, definition)
		}
	}

	partition := ""

	if t.Partition != nil {
		partition = fmt.Sprintf(" PARTITION BY %s (%s)", t.Partition.By, strings.Join(t.Partition.Columns, ", "))
	}

	foreignKeys := ""

	for _, fk := range t.ForeignKeys {
		foreignKeys += fk.Add()
	}

	indexes := ""

	for _, idx := range t.Indexes {
		indexes += idx.Create()
	}

	return fmt.Sprintf(
		"CREATE TABLE %s (%s)%s;%s%s",
		t.Name,
		strings.TrimRight(columns, ","),
		partition,
		foreignKeys,
		indexes)
}

// Drop generate sql for deleting the table
func (t *Table) Drop() string {
	foreignKeys := ""

	for _, fk := range t.ForeignKeys {
		foreignKeys += fk.Drop()
	}

	indexes := ""

	for _, idx := range t.Indexes {
		indexes += idx.Drop()
	}

	return fmt.Sprintf("%s%sDROP TABLE %s;", foreignKeys, indexes, t.Name)
}
