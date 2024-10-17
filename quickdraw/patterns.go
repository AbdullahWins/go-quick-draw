package quickdraw

import (
	"fmt"
	"sort"
)

// FindPatterns finds and counts patterns in the given numbers
func FindPatterns(numbers []string, minMatchedCount int) {
	slotLength := 2 // Each slot is 2 digits
	slotCount := 7  // Total slots for each number
	slotFrequency := make(map[string]int)

	// Count occurrences of each slot pattern
	for _, num := range numbers {
		numLength := len(num)
		for i := 1; i <= slotCount; i++ { // Checking up to 7 slots
			patternEnd := i * slotLength
			if patternEnd <= numLength { // Check if the slice is within bounds
				pattern := num[:patternEnd]
				slotFrequency[pattern]++
			}
		}
	}

	// Create a slice to hold patterns and their counts for sorting
	type PatternCount struct {
		Pattern string
		Count   int
	}

	var patterns []PatternCount
	for pattern, count := range slotFrequency {
		if len(pattern) > 4 && count >= minMatchedCount { // Only consider patterns longer than 4 digits (2 slots) and meeting min count
			patterns = append(patterns, PatternCount{Pattern: pattern, Count: count})
		}
	}

	// Sort patterns by length of the pattern (descending) and count (descending)
	sort.Slice(patterns, func(i, j int) bool {
		if len(patterns[i].Pattern) != len(patterns[j].Pattern) {
			return len(patterns[i].Pattern) > len(patterns[j].Pattern) // Sort by length
		}
		return patterns[i].Count > patterns[j].Count // Then sort by count
	})

	// Output the sorted patterns
	for _, pc := range patterns {
		fmt.Printf("Pattern \"%s\" occurs %d times.\n", pc.Pattern, pc.Count)
	}
}
