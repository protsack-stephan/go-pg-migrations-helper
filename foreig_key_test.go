package pgmigrations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const fkTestTable = "address"
const fkTestProperty = "person_id"
const fkTestReferenceTable = "people"
const fkTestReferenceProperty = "id"
const fkTestOnDeleteAction = ActionCascade
const fkTestAddSQL = "ALTER TABLE %s ADD CONSTRAINT fk_%s_%s_ref_%s_%s FOREIGN KEY (%s) REFERENCES %s (%s) ON DELETE %s;"
const fkTestDropSQL = "ALTER TABLE %s DROP CONSTRAINT fk_%s_%s_ref_%s_%s;"

func TestForeignKey(t *testing.T) {
	fk := ForeignKey{
		Table:               fkTestTable,
		Properties:          []string{fkTestProperty},
		ReferenceTable:      fkTestReferenceTable,
		ReferenceProperties: []string{fkTestReferenceProperty},
		OnDelete:            fkTestOnDeleteAction,
	}

	assert.Equal(t,
		fmt.Sprintf(
			fkTestAddSQL,
			fkTestTable,
			fkTestTable,
			fkTestProperty,
			fkTestReferenceTable,
			fkTestReferenceProperty,
			fkTestProperty,
			fkTestReferenceTable,
			fkTestReferenceProperty,
			fkTestOnDeleteAction),
		fk.Add())
	assert.Equal(t,
		fmt.Sprintf(
			fkTestDropSQL,
			fkTestTable,
			fkTestTable,
			fkTestProperty,
			fkTestReferenceTable,
			fkTestReferenceProperty),
		fk.Drop())
}
