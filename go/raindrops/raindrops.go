package raindrops

import (
	"fmt"
)

func isPrime(number int) bool { // return list
	for i := 2; i < number; i++ {
		if number%i == 0 && i != number {
			return false // add i to List
		}
	}
	return true
}

func Convert(number int) (output string) {
	fmt.Println("is", 52, "a prime?", isPrime(52))
	return ""
}

// need to get prime factors, not just check for prime.
// take number 15
// see if it divides by 3
// 15 % 3 == 0 - yes (add 3 to List)
// 15 - 5 = 10
// 10 % 3 == false
// 10 % 5 == 0 - yes (add 5 to List)
// 10 - 5 = 5
// 5 % 3 && 5 != 5 = false
// 5 % 5 && 5 != 5 = false
// 5 % 7 = false
// final List of factors of prime [3,5]
