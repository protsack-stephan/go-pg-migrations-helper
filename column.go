package pgmigrations

// Column struct to represent column in database
type Column struct {
	Table string
	Name  string
	Type  string
}

// Add creating new column inside the table
func (column *Column) Add() string {
	return "ALTER TABLE " + column.Table + " ADD COLUMN " + column.Define() + ";"
}

// Drop deleting column from table
func (column *Column) Drop() string {
	return "ALTER TABLE " + column.Table + " DROP COLUMN " + column.Name + ";"
}

// Alter update column inside the table
func (column *Column) Alter(sql string) string {
	return "ALTER TABLE " + column.Table + " ALTER COLUMN " + column.Name + " " + sql + ";"
}

// Define describing column name and type
func (column *Column) Define() string {
	return column.Name + " " + column.Type
}
