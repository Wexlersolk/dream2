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
	// Open the file in read/write mode
	file, err := os.OpenFile(writeFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read existing dreams
	var existingDreams []Dream
	if err := json.NewDecoder(file).Decode(&existingDreams); err != nil && err != io.EOF {
		return err // Handle the EOF error separately
	}

	// Concatenate new and existing dreams
	updatedDreams := concatenateDreams(newDreams, existingDreams)

	// Seek to the beginning of the file to overwrite it
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	// Truncate the file to remove old data
	if err := file.Truncate(0); err != nil {
		return err
	}

	// Sort the updated dreams
	sortDreams(updatedDreams)

	// Write the updated dreams back to the file
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

