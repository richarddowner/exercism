package octal

import (
	"fmt"
	"strconv"
	"strings"
)

func ToDecimal(octal string) (decimal int64) {
	// this is dec to octal.. not other way around
	octalint, _ := strconv.Atoi(octal)
	remainder := make([]string, 0)

	for octalint > 0 {
		remainder = append(remainder, strconv.Itoa(octalint%8))
		octalint = octalint / 8
	}

	// join array
	octal = strings.Join(remainder, "")
	fmt.Println("octal: ", octal)
	// convert to int
	blah, _ := strconv.Atoi(octal)

	return int64(blah)
}
