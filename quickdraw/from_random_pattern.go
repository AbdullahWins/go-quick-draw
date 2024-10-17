package quickdraw

import (
	"fmt"
	"sort"
	"strings"
)

// MatchResult holds the slot pattern and its corresponding match count.
type MatchResult struct {
	Pattern string
	Count   int
}

// NumberMatch holds the number and the total matched slots count.
type NumberMatch struct {
	Number       string
	TotalMatches int
}

// splitIntoSlots splits a number string into 2-digit slots.
func splitIntoSlots(number string) []string {
	var slots []string
	for i := 0; i < len(number); i += 2 {
		if i+2 <= len(number) {
			slots = append(slots, number[i:i+2])
		}
	}
	return slots
}

// countMatches checks how many slots match from the given number.
func countMatches(numbers []string, targetNumber string) ([]MatchResult, map[string]int) {
	matches := []MatchResult{}                // Slice to hold slots and their counts
	numberMatchCounts := make(map[string]int) // Map to hold the number of matches for each number

	// Split the target number into slots
	targetSlots := splitIntoSlots(targetNumber)

	// Create a map to track occurrences of each slot in the provided numbers
	slotCounts := make(map[string]int)

	// Loop through all the numbers and count the slots
	for _, num := range numbers {
		numSlots := splitIntoSlots(num)
		for _, slot := range numSlots {
			slotCounts[slot]++
		}
	}

	// Check each slot from the target number against the counted slots
	for _, slot := range targetSlots {
		if count, exists := slotCounts[slot]; exists && count >= 2 {
			matches = append(matches, MatchResult{Pattern: slot, Count: count})
			for _, num := range numbers {
				if strings.Contains(num, slot) {
					numberMatchCounts[num]++ // Increment match count for this number
				}
			}
		}
	}

	return matches, numberMatchCounts
}

// FromRandomPatterns allows you to input a number to find how many matches it has.
func FromRandomPatterns(numbers []string, targetNumber string, minMatchedSlots int) {
	// Count matching slots
	slotPatternMatches, numberMatchCounts := countMatches(numbers, targetNumber)

	// Output the sorted results for matched slots
	for _, match := range slotPatternMatches {
		fmt.Printf("Slot \"%s\" matches with %d numbers.\n", match.Pattern, match.Count)
	}

	// Count the numbers with the most matched slots
	numberMatches := make([]NumberMatch, 0, len(numberMatchCounts))
	for num, count := range numberMatchCounts {
		if count >= minMatchedSlots { // Check if the count meets the minimum requirement
			numberMatches = append(numberMatches, NumberMatch{Number: num, TotalMatches: count})
		}
	}

	// Sort the numberMatches by TotalMatches in descending order
	sort.Slice(numberMatches, func(i, j int) bool {
		return numberMatches[i].TotalMatches > numberMatches[j].TotalMatches
	})

	// Output the numbers with the most matched slots
	fmt.Printf("\nNumbers with %d or more slots matched:\n", minMatchedSlots)
	for _, nm := range numberMatches {
		fmt.Printf("Number \"%s\" has %d matched slots.\n", nm.Number, nm.TotalMatches)
	}
}
