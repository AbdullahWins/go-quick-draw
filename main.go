package main

import (
	"fmt"

	"github.com/abdullahwins/go-quick-draw/quickdraw"
)

func main() {
	// Get the generated strings
	generatedNumbers := quickdraw.GenerateStrings()
	fmt.Println("generatedNumbers:", generatedNumbers)

	quickdraw.FindPatterns()

	quickdraw.SelectWinner()

	// Wait for user input before closing
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
