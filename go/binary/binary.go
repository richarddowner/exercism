package binary

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func ToDecimal(binary string) (decimal int) {

	// loop over string backwards

	for i := 0; i < utf8.RuneCountInString(binary); i++ {
		if binary[i] == '1' {
			fmt.Println("its a 1")
			decimal = decimal + math.Pow(float64(2), float64(i))
			fmt.Println("decimal: ", decimal)
		}
	}

	return
}
