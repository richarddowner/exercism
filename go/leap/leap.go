package leap

// If year is divisible by 400 then leap year
// else if year is divisible by 100 then not leap year
// else if year is divisible by 4 then leap year
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0)
}
