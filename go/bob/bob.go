package bob

import (
	"strings"
)

func Hey(words string) string {

	if words == strings.ToUpper(words) && strings.ToLower(words) != strings.ToUpper(words) {
		return "Woah, chill out!"
	}

	if strings.HasSuffix(words, "?") {
		return "Sure."
	}

	return "Whatever."
}
