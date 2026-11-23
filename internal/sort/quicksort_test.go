package sort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
	}{
		{"nil", nil},
		{"empty", []int{}},
		{"single", []int{1}},
		{"two", []int{2, 1}},
		{"sorted", []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}},
		{"duplicates", []int{3, 1, 3, 2, 1, 3, 2}},
		{"all equal", []int{7, 7, 7, 7, 7}},
		{"negative", []int{0, -3, 8, -1, 4, -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Clone(tc.in)
			QuickSort(got)

			want := slices.Clone(tc.in)
			slices.Sort(want)

			if !slices.Equal(got, want) {
				t.Fatalf("QuickSort(%v) = %v, want %v", tc.in, got, want)
			}
		})
	}
}

func TestQuickSortStrings(t *testing.T) {
	got := []string{"pear", "apple", "fig", "banana", "apple"}
	QuickSort(got)

	want := []string{"apple", "apple", "banana", "fig", "pear"}
	if !slices.Equal(got, want) {
		t.Fatalf("QuickSort = %v, want %v", got, want)
	}
}

func TestQuickSortRandom(t *testing.T) {
	for size := 0; size <= 500; size++ {
		in := make([]int, size)
		for i := range in {
			in[i] = rand.Intn(100)
		}

		got := slices.Clone(in)
		QuickSort(got)

		want := slices.Clone(in)
		slices.Sort(want)

		if !slices.Equal(got, want) {
			t.Fatalf("size %d: QuickSort(%v) = %v, want %v", size, in, got, want)
		}
	}
}
