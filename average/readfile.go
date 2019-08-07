package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// read a line from the file
	for scanner.Scan() {
		// print the line
		fmt.Println(scanner.Text())
	}

	// close the file to free resources
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// if there was an error scanning the file, report it and exit.
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
