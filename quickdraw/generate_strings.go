package quickdraw

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Function to generate a random number within a specific range (1 to 30)
func generateRandomNumber(min, max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0, err // Return error if random generation fails
	}
	return n.Int64() + min, nil // Return the generated number
}

// Function to generate a single entry with 14 digits in the specified format
func generateEntry(min, max int64, slots int) string {
	entry := ""
	for i := 0; i < slots; i++ { // Generate a number for each slot
		number, err := generateRandomNumber(min, max)
		if err != nil {
			panic(err) // Handle error if needed
		}
		entry += fmt.Sprintf("%02d", number) // Append the number formatted as two digits
	}
	return entry
}

// Function to generate multiple entries
func generateMultipleEntries(min, max int64, slots, count int) []string {
	entries := make([]string, count)
	for i := 0; i < count; i++ {
		entries[i] = generateEntry(min, max, slots)
	}
	return entries
}

func GenerateStrings() {
	// Set the range for random numbers (1 to 30 for the slots)
	min := int64(1)
	max := int64(30)

	// Define how many entries you want to generate
	numberOfEntries := 100

	// Generate multiple entries with exactly 14 digits (7 slots of 2 digits each)
	slots := 7
	randomEntries := generateMultipleEntries(min, max, slots, numberOfEntries)

	// Print the generated entries in the desired format
	fmt.Println("Generated Entries:")
	for i, entry := range randomEntries {
		if i < len(randomEntries)-1 {
			fmt.Printf("\"%s\", ", entry) // Print each entry within double quotes, followed by a comma
		} else {
			fmt.Printf("\"%s\"", entry) // Last entry without a trailing comma
		}
	}
}
