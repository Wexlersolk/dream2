package main

import (
	"encoding/json"
	"os"
)

func readFile(readFilePath string) ([]Dream, error) {
	file, err := os.Open(readFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return decodeDreams(file)
}

func writeFile(writeFilePath string, dreams []Dream) error {
	file, err := os.Create(writeFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // Set indentation for better readability
	if err := encoder.Encode(dreams); err != nil {
		return err
	}

	return nil
}
