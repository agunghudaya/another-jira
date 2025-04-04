package utils

import (
	"errors"
	"time"
)

// ParseStringToTime parses a string into a time.Time object based on the provided layout.
func ParseStringToTime(dateStr string, layout string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date string is empty")
	}

	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
