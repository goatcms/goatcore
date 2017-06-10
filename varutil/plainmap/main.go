// Package plainmap provides primitives for convert types to plainmap.
// Plain map is one level key map. It contains keys like "lvl1.lvl2".
package plainmap

import "strings"

// Any represent any type
type Any interface{}

func formatStringJSON(s string) string {
	return "\"" + strings.Replace(s, "\"", "\\\"", -1) + "\""
}
