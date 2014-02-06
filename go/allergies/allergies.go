/*
Package allergies implements
Allergies: 	a function that takes an
			allergy score and returns
			a list of allergy matches.
AllergicTo: a function that takes an
			allergy score and an allergy
			name and returns true if
			the score matches the name.
*/
package allergies

import (
	"strconv"
)

var allergiesList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score int) []string {
	allergyMatches := make([]string, 0)
	byte := strconv.FormatInt(int64(score), 2)
	position := 0
	for bit := len(byte) - 1; bit >= 0; bit-- {
		if byte[bit] == '1' && position < len(allergiesList) {
			allergyMatches = append(allergyMatches, allergiesList[position])
		}
		position++
	}
	return allergyMatches
}

func AllergicTo(score int, name string) bool {
	for _, allergy := range Allergies(score) {
		if allergy == name {
			return true
		}
	}
	return false
}
