package arrays_and_slices

import "fmt"

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, slice := range numbersToSum {
		// fmt.Println(slice)

		// sum := int[]
		sum := []int{}

		total := 0

		for _, number := range slice {
			// fmt.Println(number)
			total += number
		}

		sum = append(sum, total)

		sums = append(sums, sum)
	}

	fmt.Println(sums)

	return
}
