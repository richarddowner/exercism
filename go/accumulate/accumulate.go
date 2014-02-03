/*
Package accumulate implements Accumulate;
a function to map a given fucntion to a
given string.
*/
package accumulate

func Accumulate(input []string, mutate func(string) string) []string {
	output := make([]string, len(input))
	for i, item := range input {
		output[i] = mutate(item)
	}
	return output
}
