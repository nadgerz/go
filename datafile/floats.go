// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([3]float64, error) {
	// Declare the array we'll be returning
	var numbers [3]float64

	// open the provided file for reading
	file, err := os.Open(fileName)
	if err != nil {
		// If there was an error opening the file, return it.
		return numbers, err
	}

	// This variable will track which array index we should assign to
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert the file line string to a float64
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)

		// If there was an errro converting the line to a number, return it
		if err != nil {
			return numbers, err
		}

		// Move to next array index.
		i++
	}

	// close the file to free resources
	err = file.Close()
	if err != nil {
		return numbers, err
	}

	// if there was an error scanning the file, report it and exit.
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	// If we go this far there were no errors, so
	// return the array of numbers and a 'nil' error
	return numbers, nil
}
