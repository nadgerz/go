package arrays_and_slices

import (
	"fmt"
	"reflect"
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
func TestSumAll(t *testing.T) {

	t.Run("new slice with totals for two slices passed in", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

// total all but the head value across multiple slices...
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("total all but the head value for a variable number of slices passed in", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum when there are empty arrays", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{5, 9})
		want := []int{0, 9}

		// checkSums(t, got, want)
		checkSums(t, got, "dave")
	})
}

//
//    EXAMPLES
//
func ExampleSumAll1() {
	total := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Printf("%#v\n", total)
	// Output: []int{3, 9}
}

func ExampleSumAll2() {
	total := SumAll([]int{1, 1, 1})
	fmt.Printf("%#v\n", total)
	// Output: []int{3}
}

func ExampleSumAllTails1() {
	total := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Printf("%#v\n", total)
	// Output: []int{2, 9}
}

func ExampleSumAllTails2() {
	total := SumAllTails([]int{1, 1, 1})
	fmt.Printf("%#v\n", total)
	// Output: []int{2}
}
