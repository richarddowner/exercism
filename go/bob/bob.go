package bob

import (
	"strings"
)

func Hey(words string) string {

	// and must contain at least 1 uppercase char
	if words == strings.ToUpper(words) {
		return "Woah, chill out!"
	}

	if strings.HasSuffix(words, "?") {
		return "Sure."
	}

	return "Whatever."
}
