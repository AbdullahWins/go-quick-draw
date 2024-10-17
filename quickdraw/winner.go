package quickdraw

import (
	"math/rand"
	"time"
)

// Determine the winner number based on the matching criteria
func DetermineWinner(tickets []string, fromStartCount int, fromEndCount int, randomCount int) string {
	if len(tickets) == 0 {
		return "No tickets available"
	}

	var bestMatch string
	bestMatchScore := -1 // Initialize with a low score

	// Iterate through each ticket to find the best match
	for _, ticket := range tickets {
		// Calculate the score based on the matching criteria
		matchScore := calculateMatchScore(ticket, fromStartCount, fromEndCount, randomCount)
		if matchScore > bestMatchScore {
			bestMatchScore = matchScore
			bestMatch = ticket
		}
	}

	return bestMatch
}

// Match digits from start, end, and random slots and return a score
func calculateMatchScore(ticket string, fromStartCount, fromEndCount, randomCount int) int {
	originalDigits := []rune(ticket)
	matchedIndexes := make(map[int]bool)
	score := 0

	// Match from the start
	for i := 0; i < fromStartCount && i < len(originalDigits); i++ {
		matchedIndexes[i] = true
		score++ // Increment score for each matched slot from start
	}

	// Match from the end
	for i := 0; i < fromEndCount && i < len(originalDigits); i++ {
		endIndex := len(originalDigits) - 1 - i
		matchedIndexes[endIndex] = true
		score++ // Increment score for each matched slot from end
	}

	// Create a new random generator with a custom source
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Match random slots
	randomMatches := 0
	for randomMatches < randomCount {
		randomIndex := rng.Intn(len(originalDigits))

		// Ensure we're not matching the same digit more than once
		if !matchedIndexes[randomIndex] {
			matchedIndexes[randomIndex] = true
			randomMatches++
			score++ // Increment score for each random match
		}
	}

	// Return the total score for the matched slots
	return score
}
