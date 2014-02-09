package atbash

import (
	"regexp"
	"strings"
)

func Atbash(plaintext string) (ciphertext string) {
	cipher := "zyxwvutsrqponmlkjihgfedcba0123456789"
	cipherkey := "abcdefghijklmnopqrstuvwxyz0123456789"

	// remove punctuation
	plaintext = strings.Join(regexp.MustCompile("([a-zA-Z0-9])+").FindAllString(plaintext, -1), "")

	// convert to lower and remove spaces
	plaintext = strings.Replace(strings.ToLower(plaintext), " ", "", -1)

	for i := 0; i < len(plaintext); i++ {
		if i != 0 && i%5 == 0 {
			ciphertext += " "
		}
		ciphertext += string(cipherkey[strings.Index(cipher, string(plaintext[i]))])
	}
	return
}
