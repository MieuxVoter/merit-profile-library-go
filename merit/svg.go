package merit

// A bunch of SVG generation utilities.

import (
	"regexp"
	"strings"
)

var spacesRegex = regexp.MustCompile("\\s+")
var lettersRegex = regexp.MustCompile("[a-zA-Z]\\s+")

func trimPathWhitespaces(path string) string {
	path = spacesRegex.ReplaceAllString(path, " ")
	path = lettersRegex.ReplaceAllStringFunc(path, func(s string) string { return s[:1] })
	path = strings.TrimSpace(path)
	return path
}
