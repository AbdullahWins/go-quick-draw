package quickdraw

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
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

// GenerateStrings generates random entries and returns them in a formatted string
func GenerateStrings() string {
	// Set the range for random numbers (1 to 30 for the slots)
	min := int64(1)
	max := int64(30)

	// Define how many entries you want to generate
	numberOfEntries := 10000

	// Generate multiple entries with exactly 14 digits (7 slots of 2 digits each)
	slots := 7
	randomEntries := generateMultipleEntries(min, max, slots, numberOfEntries)

	// Prepare the output in the desired format
	var sb strings.Builder
	sb.WriteString("{")

	for i, entry := range randomEntries {
		if i < len(randomEntries)-1 {
			sb.WriteString(fmt.Sprintf("\"%s\", ", entry)) // Append each entry with a comma
		} else {
			sb.WriteString(fmt.Sprintf("\"%s\"", entry)) // Last entry without a trailing comma
		}
	}
	sb.WriteString("}")

	// Return the generated entries in the desired format
	return sb.String()
}
