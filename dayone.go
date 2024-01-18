package main

import (
	"fmt"
	"os"
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

func getCalibration(calibrationDoc string) int {
	// List of number characters
	numChars := "1234567890"

	var total int = 0

	// Separate doc by new line
	var lines = strings.Split(calibrationDoc, "\n")

	for _, line := range lines {
		var numberBuilder strings.Builder

		// Find char of first and last number in string
		firstNumChar := line[strings.IndexAny(line, numChars)]
		lastNumChar := line[strings.LastIndexAny(line, numChars)]

		numberBuilder.WriteByte(firstNumChar)
		numberBuilder.WriteByte(lastNumChar)

		// Convert number builder to integer
		num, _ := strconv.Atoi(numberBuilder.String())

		// Add num to total
		total += num
	}

	return total
}

func main() {
	calibrationDoc := getFileContents("dayone.txt")
	calibrationNum := getCalibration(calibrationDoc)

	fmt.Println("Calibration Number:", calibrationNum)
}
