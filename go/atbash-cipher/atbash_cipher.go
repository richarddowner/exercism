package atbash

import (
	// "regexp"
	// "fmt"
	"strings"
)

func Atbash(plaintext string) (ciphertext string) {
	cipher := "zyxwvutsrqponmlkjihgfedcba0123456789"
	cipherkey := "abcdefghijklmnopqrstuvwxyz0123456789"

	// remove punctuation - there must be a better way, this is horrible
	// tmpPlainText := ""
	// tmp := regexp.MustCompile("([a-zA-Z0-9])+").FindAllString(plaintext, -1)
	// for _, char := range tmp {
	// 	tmpPlainText += char
	// }
	// plaintext = tmpPlainText

	// convert to lower and remove spaces

	plaintext = strings.ToLower(plaintext)
	plaintext = strings.Replace(plaintext, " ", "", -1)

	for i := 0; i < len(plaintext); i++ {

		// add a space every 5 chars
		if i != 0 && i%5 == 0 {
			ciphertext += " "
		}

		character := string(plaintext[i])
		index := strings.Index(cipher, character)

		if index >= 0 {
			ciphertext += string(cipherkey[index])
		}

	}

	return strings.TrimRight(ciphertext, " ")
}
