package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberToSum ...[]int) (sums []int) {
	return
}
