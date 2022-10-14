package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, e := time.Parse("1/2/2006 15:04:05", date)
	if e == nil {
		return t
	}
	t, e = time.Parse("January 2, 2006 15:04:05", date)
	if e == nil {
		return t
	}
	t, _ = time.Parse("Monday, January 2, 2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := time.Now()
	s := Schedule(date)
	if s.Year() < t.Year() {
		return true
	}
	return t.Unix() > s.Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date).Hour()
	return hour > 11 && hour < 19
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule(fmt.Sprintf("9/15/%d 00:00:00", time.Now().Year()))
}
