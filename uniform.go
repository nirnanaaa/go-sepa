package sepa

import (
	"regexp"
	"strings"
)

const emptyStringRegex = `\s+`

// UniformString formats a string, removing whitespaces and capitalizing it.
func UniformString(input string) string {
	re := regexp.MustCompile(emptyStringRegex)
	sanitized := re.ReplaceAllString(input, "")
	return strings.ToUpper(sanitized)
}
