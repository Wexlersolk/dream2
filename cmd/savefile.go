package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func readFile(readFilePath string) ([]Dream, error) {
	file, err := os.Open(readFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return decodeDreams(file)
}

func writeFile(writeFilePath string, newDreams []Dream) error {

	file, err := os.OpenFile(writeFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var existingDreams []Dream
	if err := json.NewDecoder(file).Decode(&existingDreams); err != nil && err != io.EOF {
		return err
	}

	updatedDreams := concatenateDreams(newDreams, existingDreams)

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	sortDreams(updatedDreams)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(updatedDreams); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func sortDreams(newDreams []Dream) {
	sort.Slice(newDreams, func(i, j int) bool {
		return newDreams[i].Date.Before(newDreams[j].Date)
	})
}

func concatenateDreams(newDreams []Dream, existingDreams []Dream) []Dream {
	dreamMap := make(map[time.Time]bool)

	for _, newDream := range newDreams {
		dreamMap[newDream.Date] = true
	}

	var updatedDreams []Dream

	for _, existingDream := range existingDreams {
		if _, exists := dreamMap[existingDream.Date]; !exists {
			updatedDreams = append(updatedDreams, existingDream)
		}
	}

	updatedDreams = append(updatedDreams, newDreams...)
	return updatedDreams
}
