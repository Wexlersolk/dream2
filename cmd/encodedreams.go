package main

import (
	"bufio"
	"fmt"
	"os"
)

func encodeDreams(dreams []Dream, filename string) error {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, dream := range dreams {

		weekdayStr := encodeWeekday(dream.Weekday)
		dateStr := encodeDate(dream.Date)
		scoreStr, err := encodeScore(dream.Score)
		if err != nil {
			return fmt.Errorf("invalid score: %w", err)
		}

		header := fmt.Sprintf("==%s== *%s* Score: %s\n", weekdayStr, dateStr, scoreStr)
		if _, err := writer.WriteString(header); err != nil {
			return fmt.Errorf("failed to write dream header: %w", err)
		}

		for taskDesc, completed := range dream.Tasks {
			taskLine := encodeTask(completed, taskDesc)
			if _, err := writer.WriteString(taskLine + "\n"); err != nil {
				return fmt.Errorf("failed to write task: %w", err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writes: %w", err)
	}

	return nil
}
