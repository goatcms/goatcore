package termformatter

import "strings"

const (
	maxLineLength = 320
)

var (
	emptyLine = strings.Repeat(" ", maxLineLength)
)

type FormatLineCB func(words []string, lineMax int) string
