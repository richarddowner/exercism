package leap

// If year is divisible by 400 then leap year
// else if year is divisible by 100 then not leap year
// else if year is divisible by 4 then leap year

// if divisible by 4 but also divisible by 100 and not divisble by 400
func IsLeapYear(year int) bool {
	/*
		if year%400 == 0 {
			return true
		} else if year%100 == 0 {
			return false
		} else if year%4 == 0 {
			return true
		} else {
			return false
		}
	*/
	if year%400 == 0 && year%100 == 0 && year%4 == 1 {
		return false
	}
	return true
}
