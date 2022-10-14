// Gregorian calendar leap-year calculation
package leap

// true when the given year falls on a leap year in the Gregorian calendar
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	switch {
	case year%100 != 0:
		return true
	case year%100 == 0 && year%400 == 0:
		return true
	default:
		return false
	}
}
