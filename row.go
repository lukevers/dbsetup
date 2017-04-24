package main

import (
	"fmt"
	"strings"
)

// Row is a representation of a database row that both contains a map of where
// equal statements, and a map of column:value updates for the row.
type Row struct {
	Where  map[string]interface{}
	Update map[string]interface{}
}

// Sanitize sanitizes the data prior to being used in database queries. For
// example: any "NULL" strings are replaced with actual nil values.
func (r Row) Sanitize(i map[string]interface{}) map[string]interface{} {
	for k, v := range i {
		if v == "NULL" {
			i[k] = nil
		}
	}

	return i
}

// String converts a map[string]interface{} to a human readable string.
func (r Row) String(i map[string]interface{}) string {
	str := ""
	for k, v := range i {
		str = fmt.Sprintf(
			"%s, %s=%s",
			str,
			k,
			v,
		)
	}

	return fmt.Sprintf(
		"(%s )",
		strings.Trim(str, ","),
	)
}
