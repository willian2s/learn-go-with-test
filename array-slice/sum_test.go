package arrayslice

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum collection of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{5, 4})
		want := []int{3, 9}

		checkSumAll(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{5, 4})
		want := []int{2, 4}

		checkSumAll(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSumAll(t, got, want)
	})
}

func checkSumAll(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5}
		Sum(numbers)
	}
}
