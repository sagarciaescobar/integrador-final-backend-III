package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func StringToDate(s string) (time.Time, error) {
	regex := "^\\d{4}-\\d{2}-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
	match, _ := regexp.MatchString(regex, s)
	if match {
		formated := strings.Replace(s, " ", "T", 1) + "Z"
		return time.Parse(time.RFC3339, formated)
	}
	return time.Time{}, errors.New("Invalid Format date, use YYYY-MM-DD hh:mm:ss")
}
