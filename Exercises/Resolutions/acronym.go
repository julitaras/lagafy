//Package acronym imports strings just for using in one place.
package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Abbreviate does an interesting job in abbreviating.
func Abbreviate(s string) string {
	var re = regexp.MustCompile(`[a-zA-Z]([a-z]+|[A-Z]+)`)
	var abbr string

	for _, word := range re.FindAllString(s, -1) {
		abbr += string([]rune(word)[0])
	}
	return strings.ToUpper(abbr)
}
