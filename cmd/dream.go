package main

import (
	"errors"
	"time"
)

type Dream struct {
	score int
	day   time.Time
	tasks []string
}

func NewDream(score int, day time.Time, tasks []string) (*Dream, error) {
	if score < 0 || score > 10 {
		return nil, errors.New("score must be between 0 and 10")
	}
	return &Dream{score: score, day: day, tasks: tasks}, nil
}
