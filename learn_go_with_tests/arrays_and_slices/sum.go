package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTry1(numbersToSum ...[]int) (sums []int) {

	for _, slice := range numbersToSum {

		total := 0

		for _, number := range slice {
			total += number
		}

		sums = append(sums, total)
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, slice := range numbersToSum {

		sums = append(sums, Sum(slice))
	}

	return
}
