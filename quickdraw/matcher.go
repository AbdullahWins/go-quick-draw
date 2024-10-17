package quickdraw

import (
	"fmt"
	"sort"
	"strconv"
)

// SlotMatch holds the slot pattern and its corresponding match count.
type SlotMatch struct {
	Pattern string
	Count   int
}

// countSlotMatches checks how many slots match from the start of the given number.
func countSlotMatches(numbers []string, numberToMatch string, slotLength int) []SlotMatch {
	var slotMatches []SlotMatch // Slice to hold slots and their counts

	// Loop through the numberToMatch to check increasing slot lengths
	for i := slotLength; i <= len(numberToMatch); i += slotLength {
		slotToMatch := numberToMatch[:i]
		matchCount := 0 // Reset match count for this slot length

		for _, num := range numbers {
			// Check if the current slot matches the start of any number in the list
			if len(num) >= len(slotToMatch) && num[:len(slotToMatch)] == slotToMatch {
				matchCount++ // Increment the count if a match is found
			}
		}

		// Add to the slice only if there are 2 or more matches
		if matchCount >= 2 {
			slotMatches = append(slotMatches, SlotMatch{Pattern: slotToMatch, Count: matchCount})
		}
	}

	// Sort the matches first by numeric value (descending), then by count (descending)
	sort.Slice(slotMatches, func(i, j int) bool {
		numI, _ := strconv.Atoi(slotMatches[i].Pattern) // Convert pattern to int
		numJ, _ := strconv.Atoi(slotMatches[j].Pattern) // Convert pattern to int
		if numI != numJ {
			return numI > numJ // Sort by numeric value
		}
		return slotMatches[i].Count > slotMatches[j].Count // Then sort by count
	})

	return slotMatches
}

// SelectWinner allows you to input a number to find how many matches it has.
func SelectWinner(numbers []string, numberToMatch string) {
	slotLength := 2 // Each slot is 2 digits

	// Count matching slots
	slotMatches := countSlotMatches(numbers, numberToMatch, slotLength)

	// Output the sorted results
	for _, sm := range slotMatches {
		fmt.Printf("Slot \"%s\" matches with %d numbers.\n", sm.Pattern, sm.Count)
	}
}
