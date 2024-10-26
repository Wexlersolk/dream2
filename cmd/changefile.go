package main

import (
	"errors"
)

func changefileDreams(newDreams []Dream, existingDreams []Dream, cfg config) ([]Dream, error) {
	var updatedDreams []Dream

	if len(newDreams) == 0 {
		return nil, errors.New("newDreams slice is empty")
	}
	if len(existingDreams) == 0 {
		return nil, errors.New("existingDreams slice is empty")
	}

	if cfg.Operation != Left && cfg.Operation != Right {
		return nil, errors.New("invalid operation: must be Left or Right")
	}

	if cfg.Operation == Left {
		for _, existingDream := range existingDreams {
			for _, newDream := range newDreams {
				if newDream.Date.Before(existingDream.Date) && newDream.Date.After(existingDream.Date.AddDate(0, 0, -7)) {
					updatedDreams = append(updatedDreams, newDream)
				}
			}
		}
	} else if cfg.Operation == Right {
		for _, existingDream := range existingDreams {
			for _, newDream := range newDreams {
				if newDream.Date.After(existingDream.Date) && newDream.Date.Before(existingDream.Date.AddDate(0, 0, 7)) {
					updatedDreams = append(updatedDreams, newDream)
				}
			}
		}
	}

	return updatedDreams, nil
}

