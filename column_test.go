package pgmigrations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const columnTestTable = "people"
const columnTestName = "id"
const columnTestType = "bigserial not null"
const columnTestAlter = "SET not null"
const columnTestAddSQL = "ALTER TABLE %s ADD COLUMN %s %s;"
const columnTestDefineSQL = "%s %s"
const columnTestAlterSQL = "ALTER TABLE %s ALTER COLUMN %s %s;"
const columnTestDropSQL = "ALTER TABLE %s DROP COLUMN %s;"

func TestColumn(t *testing.T) {
	column := Column{
		Table: columnTestTable,
		Name:  columnTestName,
		Type:  columnTestType,
	}

	assert.Equal(t, fmt.Sprintf(columnTestDefineSQL, columnTestName, columnTestType), column.Define())
	assert.Equal(t, fmt.Sprintf(columnTestAddSQL, columnTestTable, columnTestName, columnTestType), column.Add())
	assert.Equal(t, fmt.Sprintf(columnTestAlterSQL, columnTestTable, columnTestName, columnTestAlter), column.Alter(columnTestAlter))
	assert.Equal(t, fmt.Sprintf(columnTestDropSQL, columnTestTable, columnTestName), column.Drop())
}
