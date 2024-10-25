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
	var currentDream *Dream
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "==") && strings.Contains(line, "*") && strings.Contains(line, " Score: ") {

			dayInfo := strings.Trim(strings.TrimPrefix(line, "=="), "==")
			parts := strings.Split(dayInfo, " Score: ")
			if len(parts) != 2 {
				return nil, errors.New("invalid format: expected '==<Weekday>== *<Date>* Score: <Score>'")
			}

			dateWeekdayPart := strings.TrimSpace(parts[0])
			weekdayDateParts := strings.Split(dateWeekdayPart, " *")
			if len(weekdayDateParts) != 2 {
				return nil, errors.New("invalid format: expected '==<Weekday>== *<Date>*'")
			}

			weekdayStr := strings.Trim(strings.TrimSpace(weekdayDateParts[0]), "==")
			weekday, err := parseWeekday(weekdayStr)
			if err != nil {
				return nil, err
			}

			dateStr := strings.Trim(weekdayDateParts[1], "*")
			date, err := parseDate(dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}

			score, err := parseScore(parts[1])
			if err != nil {
				return nil, err
			}

			dream := Dream{
				Score:   score,
				Weekday: weekday,
				Date:    date,
				Tasks:   make(map[string]bool),
			}
			dreams = append(dreams, dream)
			currentDream = &dreams[len(dreams)-1]
		} else if currentDream != nil && strings.HasPrefix(line, "* [") {

			completion, task, err := parseTask(line)
			if err != nil {
				return nil, err
			}

			currentDream.Tasks[task] = completion
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dreams, nil
}
