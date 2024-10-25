package main

import (
	"log"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func drawDreams(cfg *config) error {

	dreamsToDraw, err := readFile(cfg.readFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var trimmedDreams []Dream

	for _, dream := range dreamsToDraw {
		if time.Since(dream.Date).Hours() < float64(cfg.daysToDisplay*24) {
			trimmedDreams = append(trimmedDreams, dream)
		}
	}

	dateToScoreMap := make(map[time.Time]int)
	for _, dream := range trimmedDreams {
		dateToScoreMap[dream.Date] = dream.Score
	}

	pts := make(plotter.XYs, len(dateToScoreMap))
	i := 0
	for date, score := range dateToScoreMap {
		pts[i].X = float64(date.Unix())
		pts[i].Y = float64(score)
		i++
	}

	p := plot.New()

	p.Title.Text = "Dream Scores Over Time"
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Score"

	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	p.Add(line)

	f, err := os.Create(cfg.graphFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := p.Save(6*vg.Inch, 4*vg.Inch, f.Name()); err != nil {
		return err
	}

	log.Println("Graph saved successfully as:", f.Name())
	return nil
}
