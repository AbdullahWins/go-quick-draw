package main

import (
	"github.com/abdullahwins/go-quick-draw/quickdraw"
)

func main() {
	// Generate random numbers
	numbers := quickdraw.GenerateNumbers()
	numberToMatch := "21052310050517"
	minMatchedSlots := 4

	// Find patterns in the generated numbers
	quickdraw.FindPatterns(numbers, minMatchedSlots)

	// Find how many matches the number has
	quickdraw.SelectWinner(numbers, numberToMatch, minMatchedSlots)

	// Call the function to select winning patterns
	quickdraw.SelectWinningPatterns(numbers, numberToMatch, minMatchedSlots)
}
