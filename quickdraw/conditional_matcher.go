package quickdraw

import (
	"fmt"
	"sort"
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

// countMatches checks how many slots match from the start and end of the given number.
func countMatches(numbers []string, targetNumber string, slotLength int) ([]MatchResult, map[string]int) {
	matches := []MatchResult{}                // Slice to hold slots and their counts
	numberMatchCounts := make(map[string]int) // Map to hold the number of matches for each number

	// Forward and backward match
	matches = append(matches, findMatches(numbers, targetNumber, slotLength, true, numberMatchCounts)...)
	matches = append(matches, findMatches(numbers, targetNumber, slotLength, false, numberMatchCounts)...)

	return matches, numberMatchCounts
}

// findMatches finds matches either from the front or back based on the forward flag.
func findMatches(numbers []string, targetNumber string, slotLength int, forward bool, numberMatchCounts map[string]int) []MatchResult {
	var matches []MatchResult

	// Loop through the targetNumber to check increasing slot lengths
	for i := slotLength; i <= len(targetNumber); i += slotLength {
		var slotToMatch string
		if forward {
			slotToMatch = targetNumber[:i]
		} else {
			slotToMatch = targetNumber[len(targetNumber)-i:] // Get the last 'i' characters
		}
		matchCount := 0 // Reset match count for this slot length

		for _, num := range numbers {
			// Check if the current slot matches the start or end of any number in the list
			if len(num) >= len(slotToMatch) && ((forward && num[:len(slotToMatch)] == slotToMatch) ||
				(!forward && num[len(num)-len(slotToMatch):] == slotToMatch)) {
				matchCount++             // Increment the count if a match is found
				numberMatchCounts[num]++ // Increment match count for this number
			}
		}

		// Add to the slice only if there are 2 or more matches
		if matchCount >= 2 {
			matches = append(matches, MatchResult{Pattern: slotToMatch, Count: matchCount})
		}
	}

	return matches
}

// evaluateAllMatches checks how many slots match from all starting points in the target number.
func evaluateAllMatches(numbers []string, targetNumber string, slotLength int, minMatches int) []MatchResult {
	var matchedResults []MatchResult

	// Loop through all possible starting points for slots in the target number
	for start := 0; start <= len(targetNumber)-slotLength; start++ {
		slotToMatch := targetNumber[start : start+slotLength]
		matchCount := 0

		// Check against each number in the provided list
		for _, num := range numbers {
			if len(num) >= len(slotToMatch) && num[:len(slotToMatch)] == slotToMatch {
				matchCount++
			}
		}

		// Only add to results if it meets the minimum matches criteria
		if matchCount >= minMatches {
			matchedResults = append(matchedResults, MatchResult{Pattern: slotToMatch, Count: matchCount})
		}
	}

	return matchedResults
}

// SelectWinningPatterns allows you to input a number to find how many matches it has.
func SelectWinningPatterns(numbers []string, targetNumber string, minMatchedSlots int) {
	slotLength := 2 // Each slot is 2 digits

	// Count matching slots
	slotPatternMatches, numberMatchCounts := countMatches(numbers, targetNumber, slotLength)

	// Evaluate all matches
	allSlotMatches := evaluateAllMatches(numbers, targetNumber, slotLength, minMatchedSlots)

	// Output the sorted results for forward and backward matches
	for _, match := range slotPatternMatches {
		fmt.Printf("Slot \"%s\" matches with %d numbers.\n", match.Pattern, match.Count)
	}

	// Output matches from evaluating all slots
	for _, match := range allSlotMatches {
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
