package etl

import (
	"fmt"
)

// given map[int]array (old format)
// return a map[string]int (new format)

func Transform(input map[int][]string) map[string]int {

	// for each record in the map input
	// add the key and value to the new map
	// old key becomes new value
	// old value becomes new key

	output := make(map[string]int)

	for k, v := range input {
		fmt.Println(k)
		for i, s := range v {
			fmt.Println(s[i])
		}
		fmt.Println("")
	}

	return output
}
