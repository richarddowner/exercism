package wc

func WordCount(words string) map[string]int {

	output := make(map[string]int)

	output[words] = 1

	return output
}
