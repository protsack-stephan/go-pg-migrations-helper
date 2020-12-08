package pgmigrations

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexTestColumns = []string{"id", "name"}

const indexTestTable = "people"
const indexTestMethod = IndexMethodBTree
const indexTestUnique = "CREATE UNIQUE INDEX idx_%s_%s_unique_%s ON people USING %s (%s);"
const indexTestNotUnique = "CREATE INDEX idx_%s_%s_notunique_%s ON people USING %s (%s);"
const indexTestDropUnique = "DROP INDEX idx_%s_%s_unique_%s;"
const indexTestDropNotUnique = "DROP INDEX idx_%s_%s_notunique_%s;"

func TestIndexUnique(t *testing.T) {
	index := Index{
		Table:   indexTestTable,
		Method:  IndexMethodBTree,
		Unique:  true,
		Columns: indexTestColumns,
	}

	assert.Equal(t, fmt.Sprintf(
		indexTestUnique,
		indexTestTable,
		strings.Join(indexTestColumns, "_"),
		strings.ToLower(string(indexTestMethod)),
		indexTestMethod,
		strings.Join(indexTestColumns, ", ")),
		index.Create())
	assert.Equal(t, fmt.Sprintf(
		indexTestDropUnique,
		indexTestTable,
		strings.Join(indexTestColumns, "_"),
		strings.ToLower(string(indexTestMethod))),
		index.Drop())
}

func TestIndexNotUnique(t *testing.T) {
	index := Index{
		Table:   indexTestTable,
		Method:  IndexMethodBTree,
		Unique:  false,
		Columns: indexTestColumns,
	}

	assert.Equal(t, fmt.Sprintf(
		indexTestNotUnique,
		indexTestTable,
		strings.Join(indexTestColumns, "_"),
		strings.ToLower(string(indexTestMethod)),
		indexTestMethod,
		strings.Join(indexTestColumns, ", ")),
		index.Create())
	assert.Equal(t, fmt.Sprintf(
		indexTestDropNotUnique,
		indexTestTable,
		strings.Join(indexTestColumns, "_"),
		strings.ToLower(string(indexTestMethod))),
		index.Drop())
}
