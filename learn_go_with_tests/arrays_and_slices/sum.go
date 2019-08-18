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

func SumAllTry2(numbersToSum ...[]int) (sums []int) {

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return
}

func SumAllTry3(numbersToSum ...[]int) []int {

	numberOfSlices := len(numbersToSum)
	sums := make([]int, numberOfSlices)

	for i, slice := range numbersToSum {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}
