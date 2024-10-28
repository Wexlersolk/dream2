package main

import (
	"errors"
	"fmt"
	"time"
)

func encodeWeekday(weekday time.Weekday) string {
	return weekday.String()
}

func encodeDate(date time.Time) string {
	return date.Format("02.01.2006")
}

func encodeScore(score int) (string, error) {
	if score < 0 || score > 10 {
		return "", errors.New("invalid score value")
	}
	return fmt.Sprintf("%d", score), nil
}

func encodeTask(completed bool, taskDesc string) string {
	check := "[ ]"
	if completed {
		check = "[X]"
	}
	return fmt.Sprintf("* %s | %s", check, taskDesc)
}
