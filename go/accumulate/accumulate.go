package accumulate

func Accumulate(input []string, f func(string) string) []string {
	vsm := make([]string, len(input))
	for i, v := range input {
		vsm[i] = f(v)
	}
	return vsm
}
