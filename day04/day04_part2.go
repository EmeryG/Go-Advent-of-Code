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

func getMatchCount(cardRow string) int {
	// Replace all double spaces with a single space
	cleanedString := strings.ReplaceAll(cardRow, "  ", " ")

	// Remove row prefix
	cleanedString = strings.Split(cleanedString, ": ")[1]

	// Split winning and result numbers
	parsedString := strings.Split(cleanedString, " | ")

	// Convert numbers into string slice
	winningNumbers := strings.Split(parsedString[0], " ")
	resultNumbers := strings.Split(parsedString[1], " ")

	matchCount := 0

	// Check if each result number matches winning number.
	for _, resultNumber := range resultNumbers {
		for _, winningNumber := range winningNumbers {
			if resultNumber == winningNumber {
				matchCount++
				break
			}
		}
	}

	return matchCount
}

// Might need recursion to achieve goal here
func getScratchcardCount(cardCopies map[int]int, cardMatchResults map[int]int, maxCard int) int {
	newCardCopies := make(map[int]int)
	totalCards := 0

	// Iterate through each copy of card
	for cardNum, cardCount := range cardCopies {
		// Iterate based on card results, starts at the next card
		for x := cardNum + 1; x <= cardNum+cardMatchResults[cardNum]; x++ {
			// Check if x has surpassed the max card, skip if so
			if x >= maxCard {
				break
			} else {
				_, exists := newCardCopies[x]

				// Card Count will be equal to how many card copies there were originally
				if exists {
					newCardCopies[x] += cardCount
				} else {
					newCardCopies[x] = cardCount
				}

				totalCards += cardCount
			}
		}
	}

	// Process all uncounted scratchcards recursively
	if len(newCardCopies) > 0 {
		totalCards += getScratchcardCount(newCardCopies, cardMatchResults, maxCard)
	}

	return totalCards
}

func getTotalCards(cardData string) int {
	splitRows := strings.Split(cardData, "\n")
	cardMatchResults := make(map[int]int)
	initialCardCopies := make(map[int]int)

	totalCards := 0

	// Iterate through each original card
	for cardNum, cardText := range splitRows {
		cardCount := getMatchCount(cardText)
		totalCards += 1
		cardMatchResults[cardNum] = cardCount

		if cardCount > 0 {
			initialCardCopies[cardNum] = 1
		}
	}

	// Get all additional scratch cards (recursive function)
	totalCards += getScratchcardCount(initialCardCopies, cardMatchResults, len(splitRows))

	return totalCards
}

func main() {
	cardData := getFileContents("day04.txt")
	points := getTotalCards(cardData)

	fmt.Println("Total Points", points)
}
