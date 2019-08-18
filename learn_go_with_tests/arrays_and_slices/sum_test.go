package arrays_and_slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
		}
	})
}

// take a varying number of slices, returning a new slice containing the totals for each slice passed in.
//
func TestSumAll(t *testing.T) {

	t.Run("new slice with totals for two slices passed in", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func ExampleSumAll1() {
	total := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(total)
	// Output: []int{3, 9}
}

func ExampleSumAll2() {
	total := SumAll([]int{1, 1, 1})
	fmt.Println(total)
	// Output: []int{3}
}
