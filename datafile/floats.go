// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	// Declare the slice we'll be returning
	var numbers []float64

	// open the provided file for reading
	file, err := os.Open(fileName)
	if err != nil {
		// If there was an error opening the file, return it.
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert the file line string to a float64
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		// If there was an error converting the line to a number, return it
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	// close the file to free resources
	err = file.Close()
	if err != nil {
		return nil, err
	}

	// if there was an error scanning the file, report it and exit.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	// If we go this far there were no errors, so
	// return the slice of numbers and a 'nil' error
	return numbers, nil
}
