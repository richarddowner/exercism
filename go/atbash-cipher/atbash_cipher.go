package atbash

import (
	"regexp"
	"strings"
)

func Atbash(plaintext string) (ciphertext string) {
	cipher := "zyxwvutsrqponmlkjihgfedcba0123456789"
	cipherkey := "abcdefghijklmnopqrstuvwxyz0123456789"

	// remove punctuation - there must be a better way, this is horrible
	tmpPlainText := ""
	tmp := regexp.MustCompile("([a-zA-Z0-9])+").FindAllString(plaintext, -1)
	for _, char := range tmp {
		tmpPlainText += char
	}
	plaintext = tmpPlainText

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
