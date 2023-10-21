package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type spellingBee struct {
	ExpirationUnix int      `json:"expiration"`
	DisplayWeekday string   `json:"displayWeekday"`
	DisplayDate    string   `json:"displayDate"`
	PrintDate      string   `json:"printDate"`
	CenterLetter   string   `json:"centerLetter"`
	OuterLetters   []string `json:"outerLetters"`
	ValidLetters   []string `json:"validLetters"`
	Pangrams       []string `json:"pangrams"`
	Answers        []string `json:"answers"`
	ID             int      `json:"id"`
	FreeExpiration int      `json:"freeExpiration"`
	Editor         string   `json:"editor"`
}

type pastSpellingBees struct {
	Today     spellingBee   `json:"today"`
	Yesterday spellingBee   `json:"yesterday"`
	ThisWeek  []spellingBee `json:"thisWeek"`
}

type spellingBeeHistory struct {
	Today       spellingBee      `json:"today"`
	Yesterday   spellingBee      `json:"yesterday"`
	PastPuzzles pastSpellingBees `json:"pastPuzzles"`
}

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

	spellingBeeHistory, err := extractSpellingBees(string(fileContents))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", spellingBeeHistory.Today)

	return
}

func getLatestSpellingBeeFile() (string, error) {
	files, err := os.ReadDir("./raw_spelling_bees")
	if err != nil {
		return "", err
	}
	latestFile := ""
	for _, file := range files {
		// if !strings.HasPrefix(file.Name(), "raw_spelling_bee") {
		// 	continue
		// }
		if file.Name() > latestFile {
			latestFile = file.Name()
		}
	}
	return "raw_spelling_bees/" + latestFile, nil
}

func extractSpellingBees(fileContents string) (spellingBeeHistory, error) {
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
	spellingBeeHistory := spellingBeeHistory{}
	if err := json.Unmarshal([]byte(gameData), &spellingBeeHistory); err != nil {
		return spellingBeeHistory, err
	}
	return spellingBeeHistory, nil
}
