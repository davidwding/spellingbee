package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	latestSpellingBeeFile, err := getLatestSpellingBeeFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileContents, err := os.ReadFile(latestSpellingBeeFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	spellingBees := extractSpellingBees(string(fileContents))
	_ = spellingBees

	return
}

func getLatestSpellingBeeFile() (string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}
	latestFile := ""
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "raw_spelling_bee") {
			continue
		}
		if file.Name() > latestFile {
			latestFile = file.Name()
		}
	}
	return latestFile, nil
}

func extractSpellingBees(fileContents string) string {
	windowGameData := fileContents[strings.Index(fileContents, "window.gameData"):]
	gameDataAndSuffix := windowGameData[strings.Index(windowGameData, "{"):]

	endIndex := 0
	braceBalance := 0
	for i, c := range gameDataAndSuffix {
		if c == '{' {
			braceBalance++
		} else if c == '}' {
			braceBalance--
		}

		if braceBalance == 0 {
			endIndex = i + 1
			break
		}
	}

	gameData := gameDataAndSuffix[:endIndex]

	fmt.Println(gameData)

	return ""
}
