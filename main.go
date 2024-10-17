package main

import (
	// "fmt"
	"log"

	"github.com/abdullahwins/go-quick-draw/quickdraw"
)

func main() {
	numberToMatch := "19101627121315"
	// maxOccurrences := 2
	minMatchedSlots := 6
	// generations := 300000
	filename := "numbers.txt"

	// Generate random numbers
	// numbers := quickdraw.GenerateNumbers(generations)
	// if err := quickdraw.GenerateAndSaveNumbers(numbers, filename); err != nil {
	// 	log.Printf("Error saving numbers: %v\n", err)
	// }

	// Read numbers from the file
	readNumbers, err := quickdraw.ReadNumbersFromFile(filename)
	if err != nil {
		log.Printf("Error reading numbers from file: %v\n", err)
		return
	}

	// Find the highest used slots
	// quickdraw.HighestUsedSlots(readNumbers, numberToMatch, minMatchedSlots)

	// Find the lowest used slots
	// quickdraw.LowestUsedSlots(numbers, maxOccurrences)

	// Find patterns in the generated numbers
	// quickdraw.FromStartPatterns(readNumbers, minMatchedSlots)

	// Call the function to select winning patterns
	quickdraw.CustomPattern(readNumbers, numberToMatch, minMatchedSlots)

}
