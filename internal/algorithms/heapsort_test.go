package algorithms

import (
	"slices"
	"testing"
)

func TestHeapSort(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 3, 2, 1}, []int{1, 1, 2, 3, 3}},
		{"negatives", []int{0, -3, 7, -10, 2}, []int{-10, -3, 0, 2, 7}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			HeapSort(tc.in)
			if !slices.Equal(tc.in, tc.want) {
				t.Fatalf("HeapSort = %v, want %v", tc.in, tc.want)
			}
		})
	}
}

func TestHeapSortStrings(t *testing.T) {
	in := []string{"pear", "apple", "orange", "banana"}
	HeapSort(in)
	want := []string{"apple", "banana", "orange", "pear"}
	if !slices.Equal(in, want) {
		t.Fatalf("HeapSort = %v, want %v", in, want)
	}
}
