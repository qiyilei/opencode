package quicksort

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 3, 2, 1, 2}, []int{1, 1, 2, 2, 3, 3}},
		{"negatives", []int{0, -1, 3, -5, 2}, []int{-5, -1, 0, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Sort(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}

func TestSortStrings(t *testing.T) {
	words := []string{"pear", "apple", "banana"}
	Sort(words)
	assert.Equal(t, []string{"apple", "banana", "pear"}, words)
}

func TestSortRandomized(t *testing.T) {
	for range 100 {
		input := make([]int, rand.IntN(200))
		for i := range input {
			input[i] = rand.IntN(1000) - 500
		}

		want := slices.Clone(input)
		slices.Sort(want)

		Sort(input)
		assert.Equal(t, want, input)
	}
}
