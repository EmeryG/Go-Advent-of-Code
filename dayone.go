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

// Converts first and last occurrence of each alphabetic digit from 1-9
func convertLetterDigits(str string) string {
	writtenToDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	insertNumericDigits := make(map[int]string)

	// Iterate through each written digit and store starting index in addDigits map
	for writtenDigit, numericDigit := range writtenToDigits {
		indexWritten := strings.Index(str, writtenDigit)

		if indexWritten > -1 {
			insertNumericDigits[indexWritten] = numericDigit
		}

		lastIndexWritten := strings.LastIndex(str, writtenDigit)

		if lastIndexWritten > -1 {
			insertNumericDigits[lastIndexWritten] = numericDigit
		}
	}

	var convertedStr string = str

	// Replace starting letter of written digit with numeric digit
	for index, numericDigit := range insertNumericDigits {
		convertedStr = convertedStr[:index] + numericDigit + convertedStr[index+1:]
	}

	return convertedStr
}

func getCalibration(calibrationDoc string) int {
	// List of number characters
	numChars := "1234567890"

	var total int = 0

	// Separate doc by new line
	var lines = strings.Split(calibrationDoc, "\n")

	for _, line := range lines {
		var numberBuilder strings.Builder

		// Convert alpabetic digits to numbers
		convertedLine := convertLetterDigits(line)

		// Find char of first and last number in string
		firstNumChar := convertedLine[strings.IndexAny(convertedLine, numChars)]
		lastNumChar := convertedLine[strings.LastIndexAny(convertedLine, numChars)]

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
