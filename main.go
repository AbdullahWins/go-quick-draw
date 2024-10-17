package main

import (
	"github.com/abdullahwins/go-quick-draw/quickdraw"
)

func main() {
	// Generate random numbers
	numbers := quickdraw.GenerateNumbers()

	// Find patterns in the generated numbers
	quickdraw.FindPatterns(numbers)

	// Example number to match
	numberToMatch := "21052310050517"

	// Find how many matches the number has
	quickdraw.SelectWinner(numbers, numberToMatch)
}
