package booking

import (
	"fmt"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
func Schedule(date string) time.Time {
	scheduledTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return scheduledTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	scheduledTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	if scheduledTime.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	scheduledTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if (scheduledTime.Hour() >= 12) && (scheduledTime.Hour() <= 15) {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	scheduledTime := Schedule(date)
	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %d:%d.",
		scheduledTime.Weekday(),
		scheduledTime.Month(),
		scheduledTime.Day(),
		scheduledTime.Year(),
		scheduledTime.Hour(),
		scheduledTime.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	now := time.Now()
	thisYear := time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return thisYear
}
