package main

import (
	"errors"
	"time"
)

func changefileDreams(newDreams []Dream, existingDreams []Dream, cfg config) ([]Dream, error) {
	if len(newDreams) != 7 {
		return nil, errors.New("you have not entered an appropriate number of days for a week")
	}
	if len(existingDreams) == 0 {
		return nil, errors.New("existingDreams slice is empty")
	}
	if cfg.Operation != Left && cfg.Operation != Right {
		return nil, errors.New("invalid operation: must be Left or Right")
	}

	// Determine the start and end dates for the previous or next week
	var targetStart, targetEnd time.Time
	firstDay := newDreams[0].Date

	if cfg.Operation == Left {
		targetStart = firstDay.AddDate(0, 0, -7)
		targetEnd = firstDay.AddDate(0, 0, -1)
	} else {
		targetStart = firstDay.AddDate(0, 0, 7)
		targetEnd = firstDay.AddDate(0, 0, 13)
	}

	// Filter existing dreams that fall within the target week range
	var updatedDreams []Dream
	for _, existingDream := range existingDreams {
		if (existingDream.Date.After(targetStart) || existingDream.Date.Equal(targetStart)) &&
			(existingDream.Date.Before(targetEnd) || existingDream.Date.Equal(targetEnd)) {
			updatedDreams = append(updatedDreams, existingDream)
		}
	}

	// Ensure we have exactly one week of dreams in updatedDreams
	if len(updatedDreams) != 7 {
		return nil, errors.New("could not retrieve a complete previous or next week")
	}

	return updatedDreams, nil
}
