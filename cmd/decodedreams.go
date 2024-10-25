package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func decodeDreams(file *os.File) ([]Dream, error) {
	var dreams []Dream
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "==") && strings.Contains(line, "*") && strings.Contains(line, " Score: ") {
			// Extract weekday, date, and score from the line
			dayInfo := strings.Trim(strings.TrimPrefix(line, "=="), "==") // Remove "==" from both ends
			parts := strings.Split(dayInfo, " Score: ")
			if len(parts) != 2 {
				return nil, errors.New("invalid format: expected '==<Weekday>== *<Date>* Score: <Score>'")
			}

			// Extract weekday and date
			dateWeekdayPart := strings.TrimSpace(parts[0])
			weekdayDateParts := strings.Split(dateWeekdayPart, " *")
			if len(weekdayDateParts) != 2 {
				return nil, errors.New("invalid format: expected '==<Weekday>== *<Date>*'")
			}

			// Parse weekday
			weekdayStr := strings.TrimSpace(weekdayDateParts[0])
			weekdayStr = strings.Trim(weekdayStr, "==")
			weekday, err := parseWeekday(weekdayStr)
			if err != nil {
				return nil, err
			}

			// Parse date
			dateStr := strings.Trim(weekdayDateParts[1], "*")
			date, err := parseDate(dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}

			// Parse score
			score, err := parseScore(parts[1])
			if err != nil {
				return nil, err
			}

			// Log the weekday, date, and score
			fmt.Printf("Weekday: %s, Date: %s, Score: %d\n", weekday, date.Format("2006-01-02"), score)

			// Set up the Dream struct and add it to the list
			dream := Dream{
				Score:   score,
				Weekday: weekday,
				Date:    date,
				Tasks:   make(map[string]bool),
			}
			dreams = append(dreams, dream)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dreams, nil
}
