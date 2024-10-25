package main

import (
	"errors"
	"log"
	"time"
)

type Dream struct {
	Score   int
	Weekday time.Weekday
	Date    time.Time
	Tasks   map[string]bool
}

type Operation string

const (
	Right Operation = "Right"
	Left  Operation = "Left"
	None  Operation = ""
)

type config struct {
	readFilePath  string
	writeFilePath string
	graphFilePath string
	daysToDisplay int
	Operation     ""
}

func NewDream(score int, date time.Time, tasks map[string]bool) (*Dream, error) {
	if score < 0 || score > 10 {
		return nil, errors.New("score must be between 0 and 10")
	}
	return &Dream{
		Score:   score,
		Date:    date,
		Weekday: date.Weekday(),
		Tasks:   tasks,
	}, nil
}

func run(cfg config) error {
	dreams, err := readFile(cfg.readFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = writeFile(cfg.writeFilePath, dreams)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
