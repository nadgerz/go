// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var line []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Add the line to the slice directly.
		line := scanner.Text()

		lines = append(lines, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
