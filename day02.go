package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gameValueFn func(int, string) int

func getFileContents(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

// Get results for each cube from a round
func getCubeRoundResults(round string) (int, int, int) {
	green, red, blue := 0, 0, 0

	// Checks for cube suffixes and takes number associated if matxed
	for _, cubeStr := range strings.Split(round, ", ") {
		if strings.Contains(cubeStr, " green") {
			green, _ = strconv.Atoi(strings.TrimSuffix(cubeStr, " green"))

		} else if strings.Contains(cubeStr, " red") {
			red, _ = strconv.Atoi(strings.TrimSuffix(cubeStr, " red"))

		} else if strings.Contains(cubeStr, " blue") {
			blue, _ = strconv.Atoi(strings.TrimSuffix(cubeStr, " blue"))

		}
	}

	return green, red, blue
}

// Gets game value depending on if max cubes of each color fits criteria
func getGameValueMax(gameNum int, game string) int {
	maxGreen, maxRed, maxBlue := 0, 0, 0

	// Iterate through each round and compare max
	for _, round := range strings.Split(game, "; ") {
		greenCount, redCount, blueCount := getCubeRoundResults(round)

		if greenCount > maxGreen {
			maxGreen = greenCount
		}

		if redCount > maxRed {
			maxRed = redCount
		}

		if blueCount > maxBlue {
			maxBlue = blueCount
		}
	}

	if maxGreen <= 13 && maxRed <= 12 && maxBlue <= 14 {
		return gameNum + 1
	} else {
		return 0
	}
}

func getGameValueMin(gameNum int, game string) int {
	maxGreen, maxRed, maxBlue := 0, 0, 0

	// Iterate through each round and compare max
	for _, round := range strings.Split(game, "; ") {
		greenCount, redCount, blueCount := getCubeRoundResults(round)

		if greenCount > maxGreen {
			maxGreen = greenCount
		}

		if redCount > maxRed {
			maxRed = redCount
		}

		if blueCount > maxBlue {
			maxBlue = blueCount
		}
	}

	return maxGreen * maxRed * maxBlue
}

func getGameSum(fileContents string, getGameValue gameValueFn) int {
	total := 0

	// Split day02.txt into lines
	for index, line := range strings.Split(fileContents, "\n") {
		if strings.Contains(line, "Game") {
			// Trim game prefix from string
			trimmedLine := strings.Split(line, ": ")[1]

			// Add game value to total
			total += getGameValue(index, trimmedLine)
		}
	}

	return total
}

func main() {
	fileContents := getFileContents("day02.txt")

	gameMaxSum := getGameSum(fileContents, getGameValueMax)

	fmt.Println("Game Max Sum", gameMaxSum)

	gameMinSum := getGameSum(fileContents, getGameValueMin)

	fmt.Println("Game Max Sum", gameMinSum)
}
