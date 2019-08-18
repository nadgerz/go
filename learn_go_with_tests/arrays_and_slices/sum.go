package arrays_and_slices

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

		total := 0

		for _, number := range slice {
			// fmt.Println(number)
			total += number
		}

		sums = append(sums, total)
	}

	// fmt.Println(sums)

	return
}
