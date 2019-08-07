package main

import (
	"fmt"
)

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %).2f\n", sum/sampleCount)
}
