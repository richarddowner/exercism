package leap

// if year is divisible by 4 and not 100 then is leap year
// if year is divisible by 4, 100 and 400 then is leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || (year%100 == 0 && year%400 == 0))
}
