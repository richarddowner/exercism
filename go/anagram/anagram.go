package anagram

import (
	"fmt"
	"sort"
)

// type ByLength []string
type Rune []rune

// func (s ByLength) Len() int {
// 	return len(s)
// }
func (r Rune) Len() int {
	return len(r)
}

// func (s ByLength) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }
func (r Rune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// func (s ByLength) Less(i, j int) bool {
// 	return len(s[i]) < len(s[j])
// }
func (r Rune) Less(i, j int) bool {
	return r[i] < r[j]
}

// func main() {
// 	fruits := []string{"peach", "banana", "kiwi"}
// 	sort.Sort(ByLength(fruits))
// 	fmt.Println(fruits)
// }
func stringSort(s string) string {
	runes := []rune(s)
	sort.Sort(Rune(runes))
	return string(runes)
}

// two sorted anagrams will always be the same word
func Detect(subject string, candidates []string) []string {

	var anagrams = make([]string, 0)

	fmt.Println(subject, stringSort(subject))

	return anagrams
}
