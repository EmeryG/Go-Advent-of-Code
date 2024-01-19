package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getFileContents(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func hasAdjacentSymbols(parsedSchematic []string, rowIndex int, numberIndex []int) bool {
	// Logic for which rows need to be checked
	rowRange := [2]int{rowIndex - 1, rowIndex + 2}

	if rowIndex == 0 {
		rowRange[0] = 0
	} else if rowRange[1] > len(parsedSchematic) {
		rowRange[1] = len(parsedSchematic)
	}

	// Logic for indices in the text lines that need to be checked
	lineRange := [2]int{numberIndex[0] - 1, numberIndex[1] + 1}

	if numberIndex[0] == 0 {
		lineRange[0] = 0

		// Since all lines are same length, checks length of first line
	} else if numberIndex[1] == len(parsedSchematic[0]) {
		lineRange[1] = numberIndex[1]
	}

	// Regex for any nonnumerical value that also isn't "."
	regexSymbol, _ := regexp.Compile(`[^0-9.]`)

	for _, schematicLine := range parsedSchematic[rowRange[0]:rowRange[1]] {
		// Get text adjacent to number
		adjacentText := schematicLine[lineRange[0]:lineRange[1]]

		// Check for any symbol matches
		symbolMatches := regexSymbol.FindAllStringIndex(adjacentText, -1)

		// Check if any matches populated
		if len(symbolMatches) > 0 {
			return true
		}
	}

	return false
}

func getPartSum(engineSchematic string) int {
	// Regex for any numbers
	regexNum, _ := regexp.Compile(`\d+`)

	parsedSchematic := strings.Split(engineSchematic, "\n")

	// var partNumbers []int
	partSum := 0

	for rowIndex, schematicRow := range parsedSchematic {
		// Returns 2D slice of the range of each number
		numberIndices := regexNum.FindAllStringIndex(schematicRow, -1)

		for _, numberIndex := range numberIndices {
			if hasAdjacentSymbols(parsedSchematic, rowIndex, numberIndex) {
				partNum, _ := strconv.Atoi(schematicRow[numberIndex[0]:numberIndex[1]])

				partSum += partNum
			}
		}
		break
	}

	return partSum
}

func main() {
	engineSchematic := getFileContents("day03.txt")
	partSum := getPartSum(engineSchematic)

	fmt.Println("Part Number Sum", partSum)
}
