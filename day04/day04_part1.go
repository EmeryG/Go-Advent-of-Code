package main

import (
	"fmt"
	"os"
	"strings"
)

func getFileContents(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func getRowPoints(winningNumStr, resultNumStr string) int {
	// Convert numbers into string slice
	winningNumbers := strings.Split(winningNumStr, " ")
	resultNumbers := strings.Split(resultNumStr, " ")

	total := 0

	for _, resultNumber := range resultNumbers {
		for _, winningNumber := range winningNumbers {
			if resultNumber == winningNumber {
				// If match, total gets double if above 0, or set to 1 if 0
				if total > 0 {
					total *= 2
				} else {
					total = 1
				}

				break
			}
		}
	}

	return total

}

func getCardPoints(cardData string) int {
	total := 0

	for _, row := range strings.Split(cardData, "\n") {
		// Replace all double spaces with a single space
		cleanedString := strings.ReplaceAll(row, "  ", " ")

		// Remove row prefix
		cleanedString = strings.Split(cleanedString, ": ")[1]

		// Split winning and result numbers
		parsedString := strings.Split(cleanedString, " | ")

		total += getRowPoints(parsedString[0], parsedString[1])
	}

	return total
}

func main() {
	cardData := getFileContents("day04.txt")
	points := getCardPoints(cardData)

	fmt.Println("Total Points", points)
}
