package anagram

import (
	"sort"
	"strings"
)

type Rune []rune

func (r Rune) Len() int {
	return len(r)
}

func (r Rune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rune) Less(i, j int) bool {
	return r[i] < r[j]
}

func stringSort(s string) string {
	runes := []rune(s)
	sort.Sort(Rune(runes))
	return string(runes)
}

// two sorted anagrams will always be the same word
func Detect(subject string, candidates []string) []string {

	var anagrams = make([]string, 0)

	// sort subject
	sortedSubject := stringSort(strings.ToLower(subject))

	// for each candidate
	for _, candidate := range candidates {

		candidate = strings.ToLower(candidate)

		if subject == candidate {
			continue
		}

		// sort the candidate
		sortedCandidate := stringSort(candidate)

		if sortedCandidate == sortedSubject {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
