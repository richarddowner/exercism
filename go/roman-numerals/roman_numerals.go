// Package romannumerals implements ToRomanNumeral.
package romannumerals

type kvp struct {
	k int
	v string
}

var correlations = []kvp{
	kvp{1000, "M"},
	kvp{900, "CM"},
	kvp{500, "D"},
	kvp{400, "CD"},
	kvp{100, "C"},
	kvp{90, "XC"},
	kvp{50, "L"},
	kvp{40, "XL"},
	kvp{10, "X"},
	kvp{9, "IX"},
	kvp{5, "V"},
	kvp{4, "IV"},
	kvp{1, "I"},
}

// ToRomanNumeral takes an arabic integer.
// Returns the corresponding roman numeral.
func ToRomanNumeral(arabic int) (roman string) {
	for arabic > 0 {
		for _, numeral := range correlations {
			if arabic-numeral.k >= 0 {
				arabic -= numeral.k
				roman += numeral.v
				break
			}
		}
	}
	return
}
