package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func parseWeekday(weekdayStr string) (time.Weekday, error) {
	switch weekdayStr {
	case "Monday":
		return time.Monday, nil
	case "Tuesday":
		return time.Tuesday, nil
	case "Wednesday":
		return time.Wednesday, nil
	case "Thursday":
		return time.Thursday, nil
	case "Friday":
		return time.Friday, nil
	case "Saturday":
		return time.Saturday, nil
	case "Sunday":
		return time.Sunday, nil
	default:
		return time.Sunday, fmt.Errorf("invalid weekday: %s", weekdayStr)
	}
}

// Helper function to parse the date string into a time.Time object
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("02.01.2006", dateStr)
}

// Helper function to parse the score
func parseScore(scoreStr string) (int, error) {
	var score int
	_, err := fmt.Sscanf(scoreStr, "%d", &score)
	if err != nil || score < 0 || score > 10 {
		return 0, errors.New("invalid score format")
	}
	return score, nil
}

func parseTask(taskLine string) (bool, string, error) {
	// Check for task completion status and extract description
	if !(strings.HasPrefix(taskLine, "* [ ] |") || strings.HasPrefix(taskLine, "* [X] |")) {
		return false, "", errors.New("invalid task format")
	}

	completed := strings.Contains(taskLine, "[X]")                     // True if the task is completed
	taskDesc := strings.TrimSpace(strings.SplitN(taskLine, "|", 2)[1]) // Get the task description

	return completed, taskDesc, nil
}
