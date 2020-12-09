package pgmigrations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	assert.NotNil(t, new(Partition))
}
