package pgmigrations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTableColumns = []Column{
	{
		Name: "id",
		Type: "int",
	},
	{
		Name: "name",
		Type: "varchar(255)",
	},
}

const testTableName = "spouse"
const testTablePrimaryKey = "id"
const testTableCreateResult = "CREATE TABLE %s (%s,%s,PRIMARY KEY (%s));"
const testTableDropResult = "DROP TABLE %s;"

func TestTable(t *testing.T) {
	table := Table{
		Name:    testTableName,
		Columns: testTableColumns,
		Constraints: map[Constraint][]string{
			ConstraintPrimaryKey: []string{
				Columns([]string{testTablePrimaryKey}),
			},
		},
	}

	assert.Equal(t, fmt.Sprintf(
		testTableCreateResult,
		testTableName,
		testTableColumns[0].Define(),
		testTableColumns[1].Define(),
		testTablePrimaryKey,
	),
		table.Create())
	assert.Equal(t, fmt.Sprintf(testTableDropResult, testTableName), table.Drop())
}
