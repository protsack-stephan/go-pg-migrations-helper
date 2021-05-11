package pgmigrations

import "fmt"

// Column struct to represent column in database
type Column struct {
	Table string
	Name  string
	Type  string
}

// Add creating new column inside the table
func (column *Column) Add() string {
	return fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s;", column.Table, column.Define())
}

// Drop deleting column from table
func (column *Column) Drop() string {
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", column.Table, column.Name)
}

// Alter update column inside the table
func (column *Column) Alter(sql string) string {
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s;", column.Table, column.Name, sql)
}

// Rename rename table column
func (column *Column) Rename(name string) string {
	return fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s;", column.Table, column.Name, name)
}

// Define describing column name and type
func (column *Column) Define() string {
	return fmt.Sprintf("%s %s", column.Name, column.Type)
}
