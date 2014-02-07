package romannumerals

// I : 1
// V : 5
// X : 10
// L : 50
// C : 100
// D : 500
// M : 1000

func ToRomanNumeral(arabic int) (roman string) {
	return "I"
}

// cases
// % 5 = 0
// x > y && x % y
// x < y && x % y

// 48
// tenths
// 4 % 5 = 4 ~ if 4: take (X L)
// ones
// 8 % 5 = 3 ~ (V I I I)

// 59
// tenths
// 5 % 5 = 0 ~ if 0: take (L)
// ones
// 9 % 5 = 4

// yes
// 141
// hundreths
// 1 % 5 = 1 (C)
// tenths
// 4 % 5 = 4 ~ if 4: take (X L)
// ones
// 1 % 5 = 1 (I)

// no
// 575
// hundreths
// 5 % 5 = 0 ~ if 0: take (D)
// tenths
// 7 % 5 = 2 ~ (X X)
// ones
// 5 % 5 = 0 ~ if 0: take (V)
