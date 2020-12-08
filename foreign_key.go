package pgmigrations

import (
	"fmt"
	"strings"
)

// Action postgresql action
type Action string

// ActionSetNull set fk column to null
const ActionSetNull Action = "SET NULL"

// ActionSetDefault fallback do default value
const ActionSetDefault Action = "SET DEFAULT"

// ActionRestrict restrict operation with active fk
const ActionRestrict Action = "RESTRICT"

// ActionDoNothing don't don anything
const ActionDoNothing Action = "NO ACTION"

// ActionCascade do the same thing with dependencies
const ActionCascade Action = "CASCADE"

// ForeignKey create new fk constraint
type ForeignKey struct {
	Table               string
	Properties          []string
	ReferenceTable      string
	ReferenceProperties []string
	OnDelete            Action
	OnUpdate            Action
}

// Name get foreign key name
func (fk *ForeignKey) Name() string {
	return fmt.Sprintf("fk_%s_%s_ref_%s_%s",
		fk.Table,
		strings.Join(fk.Properties, ", "),
		fk.ReferenceTable,
		strings.Join(fk.ReferenceProperties, ", "))
}

// Add alter table to add constraint
func (fk *ForeignKey) Add() string {
	onDelete := ""
	onUpdate := ""

	if len(fk.OnDelete) > 0 {
		onDelete = fmt.Sprintf(" ON DELETE %s", fk.OnDelete)
	}

	if len(fk.OnUpdate) > 0 {
		onUpdate = fmt.Sprintf(" ON UPDATE %s", fk.OnUpdate)
	}

	return fmt.Sprintf(
		"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)%s%s;",
		fk.Table,
		fk.Name(),
		strings.Join(fk.Properties, ", "),
		fk.ReferenceTable,
		strings.Join(fk.ReferenceProperties, ", "),
		onDelete,
		onUpdate)
}

// Drop generate SQL for dropping foreign key constraint
func (fk *ForeignKey) Drop() string {
	return fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s;", fk.Table, fk.Name())
}
