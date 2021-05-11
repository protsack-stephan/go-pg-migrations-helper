package pgmigrations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const columnTestTable = "people"
const columnTestName = "id"
const columnTestAlternateName = "uid"
const columnTestType = "bigserial not null"
const columnTestAlter = "SET not null"
const columnTestAddSQL = "ALTER TABLE %s ADD COLUMN %s %s;"
const columnTestDefineSQL = "%s %s"
const columnTestAlterSQL = "ALTER TABLE %s ALTER COLUMN %s %s;"
const columnTestDropSQL = "ALTER TABLE %s DROP COLUMN %s;"
const columnTestRenameSQL = "ALTER TABLE %s RENAME COLUMN %s TO %s;"

func TestColumn(t *testing.T) {
	column := Column{
		Table: columnTestTable,
		Name:  columnTestName,
		Type:  columnTestType,
	}

	assert := assert.New(t)

	assert.Equal(fmt.Sprintf(columnTestDefineSQL, columnTestName, columnTestType), column.Define())
	assert.Equal(fmt.Sprintf(columnTestAddSQL, columnTestTable, columnTestName, columnTestType), column.Add())
	assert.Equal(fmt.Sprintf(columnTestAlterSQL, columnTestTable, columnTestName, columnTestAlter), column.Alter(columnTestAlter))
	assert.Equal(fmt.Sprintf(columnTestDropSQL, columnTestTable, columnTestName), column.Drop())
	assert.Equal(fmt.Sprintf(columnTestRenameSQL, columnTestTable, columnTestName, columnTestAlternateName), column.Rename(columnTestAlternateName))
}
