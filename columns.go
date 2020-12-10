package pgmigrations

import (
	"fmt"
	"strings"
)

// Columns convert array of columns to string
func Columns(list []string) string {
	return fmt.Sprintf("(%s)", strings.Join(list, ", "))
}
