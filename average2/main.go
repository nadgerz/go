package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0

	// process each variadic argument
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func main() {
	// fmt.Println(os.Args)
	// fmt.Println(os.Args[1:])

	arguments := os.Args[1:]

	var sum float64 = 0

	for _, argument := range arguments {

		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
			// return nil, err
		}

		sum += number
	}

	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}
