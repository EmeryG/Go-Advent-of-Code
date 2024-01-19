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

func isOverlapRange(rangeA []int, rangeB []int) bool {
	return rangeA[0] <= rangeB[1] && rangeA[1] >= rangeB[0]
}

func getConnectedGears(parsedSchematic []string, rowIndex int, asteriskIndex []int) []int {
	// Logic for which rows need to be checked
	rowRange := [2]int{rowIndex - 1, rowIndex + 2}

	if rowIndex == 0 {
		rowRange[0] = 0
	} else if rowRange[1] > len(parsedSchematic) {
		rowRange[1] = len(parsedSchematic)
	}

	// Regex for any numbers
	regexNum, _ := regexp.Compile(`\d+`)

	var gears []int

	for _, schematicLine := range parsedSchematic[rowRange[0]:rowRange[1]] {
		// Find all numbers in schematicLine
		numberIndices := regexNum.FindAllStringIndex(schematicLine, -1)

		// Iterate through each number and find if it overlaps with the asterisk range
		for _, numberIndex := range numberIndices {
			if isOverlapRange(asteriskIndex, numberIndex) {
				gearNum, _ := strconv.Atoi(schematicLine[numberIndex[0]:numberIndex[1]])

				gears = append(gears, gearNum)
			}
		}
	}

	return gears
}

func getGearRatioSum(engineSchematic string) int {
	// Regex for asterisk
	regexRatio, _ := regexp.Compile(`[*]`)

	parsedSchematic := strings.Split(engineSchematic, "\n")

	// var partNumbers []int
	gearRatioSum := 0

	for rowIndex, schematicRow := range parsedSchematic {
		// Returns 2D slice of the range of each number
		asteriskIndices := regexRatio.FindAllStringIndex(schematicRow, -1)

		for _, asteriskIndex := range asteriskIndices {
			// Get gears connected to asterisks
			gears := getConnectedGears(parsedSchematic, rowIndex, asteriskIndex)

			// If 2 gears matched to asterisk, add to gear ratio sum
			if len(gears) >= 2 {
				gearRatioSum += gears[0] * gears[1]
			}
		}
	}

	return gearRatioSum
}

func main() {
	engineSchematic := getFileContents("day03.txt")
	gearRatioSum := getGearRatioSum(engineSchematic)

	fmt.Println("Gear Ratio Sum", gearRatioSum)
}
