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

func SelectWinner() {
	// Example slice of numbers
	numbers := []string{
		"21052310050527",
		"21052310050527",
		"21022310050527",
		"21052310050517",
		"21052310050527",
		"21051311050527",
		"23051503112024",
		"30222728220803",
		"21051508190208",
		"21050620052024",
	}

	// Example number to match
	numberToMatch := "21052310050517"
	slotLength := 2 // Each slot is 2 digits

	// Count matching slots
	slotMatches := countSlotMatches(numbers, numberToMatch, slotLength)

	// Output the sorted results
	for _, sm := range slotMatches {
		fmt.Printf("Slot \"%s\" matches with %d numbers.\n", sm.Pattern, sm.Count)
	}
}
