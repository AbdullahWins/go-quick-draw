package quickdraw

import (
	"fmt"
	"sort"
)

// FindPatterns finds and counts patterns in the given numbers
func LowestUsedSlots(numbers []string, maxOccurrences int, slotCount int) {
	slotLength := 2 // Each slot is 2 digits
	patternFrequency := make(map[string]int)

	// Count occurrences of each slot pattern
	for _, num := range numbers {
		numLength := len(num)
		for i := 1; i <= slotCount; i++ { // Checking up to 7 slots
			patternEnd := i * slotLength
			if patternEnd <= numLength { // Check if the slice is within bounds
				pattern := num[:patternEnd]
				patternFrequency[pattern]++
			}
		}
	}

	// Create a slice to hold patterns and their counts for sorting
	type FrequencyInfo struct {
		Pattern string
		Count   int
	}

	var filteredPatterns []FrequencyInfo
	for pattern, count := range patternFrequency {
		if len(pattern) > 4 && count <= maxOccurrences { // Only consider patterns longer than 4 digits and matching the maxOccurrences criteria
			filteredPatterns = append(filteredPatterns, FrequencyInfo{Pattern: pattern, Count: count})
		}
	}

	// Sort patterns by count (ascending)
	sort.Slice(filteredPatterns, func(i, j int) bool {
		return filteredPatterns[i].Count < filteredPatterns[j].Count // Sort by count in ascending order
	})

	// Output the sorted patterns
	for _, freqInfo := range filteredPatterns {
		fmt.Printf("Pattern \"%s\" occurs %d times.\n", freqInfo.Pattern, freqInfo.Count)
	}
}
