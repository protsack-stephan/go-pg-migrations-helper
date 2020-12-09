package pgmigrations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testColumnsList = []string{"id", "name"}

const testColumnsResult = "(id, name)"

func TestColumns(t *testing.T) {
	assert.Equal(t, testColumnsResult, Columns(testColumnsList))
}
