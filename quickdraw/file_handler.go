package quickdraw

import (
	"fmt"
	"os"
	"strings"
)

// SaveNumbersToFile saves the generated numbers to a text file
func SaveNumbersToFile(numbers []string, filename string) error {
	// Prepare the data to be saved
	data := []byte(strings.Join(numbers, "\n") + "\n") // Join numbers with newline

	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write data to the file
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

// ReadNumbersFromFile reads numbers from a text file and returns them as a slice of strings
func ReadNumbersFromFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename) // Updated to use os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Split the data into lines and return as a slice of strings
	numbers := strings.Split(string(data), "\n")
	return numbers, nil
}

// GenerateAndSaveNumbers generates numbers and saves them to a file

func GenerateAndSaveNumbers(numbers []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(numbers, "\n"))
	if err != nil {
		return err
	}
	return nil

}
