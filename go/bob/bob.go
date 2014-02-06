/*
Package bob implements Hey;
a function that allows Alice to
talk to Bob and obtain his very
limited responses.
*/
package bob

import (
	"strings"
)

func Hey(words string) string {
	if strings.Trim(words, " ") == "" {
		return "Fine. Be that way!"
	}
	if words == strings.ToUpper(words) && strings.ToLower(words) != strings.ToUpper(words) {
		return "Woah, chill out!"
	}
	if strings.HasSuffix(words, "?") {
		return "Sure."
	}
	return "Whatever."
}
