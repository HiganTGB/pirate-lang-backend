package utils

import (
	"time"
)

func isValidDateOfBirthGoTime(dob time.Time, minAge, maxAge int) bool {
	now := time.Now()

	if dob.After(now) {
		return false
	}

	years := now.Year() - dob.Year()

	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}

	if years < minAge {
		return false
	}
	if years > maxAge {
		return false
	}

	return true
}
