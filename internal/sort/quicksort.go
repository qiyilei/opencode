// Package sort provides generic sorting algorithms.
package sort

import "cmp"

// QuickSort sorts s in place in ascending order using quicksort with
// median-of-three pivot selection and a three-way partition, so slices
// containing many duplicate keys stay efficient.
func QuickSort[T cmp.Ordered](s []T) {
	quickSort(s, 0, len(s)-1)
}

func quickSort[T cmp.Ordered](s []T, lo, hi int) {
	for lo < hi {
		// Insertion sort is faster than quicksort on tiny slices.
		if hi-lo < 12 {
			insertionSort(s[lo : hi+1])
			return
		}

		lt, gt := partition(s, lo, hi)

		// Recurse into the smaller side and loop on the larger one to
		// keep the recursion depth bounded.
		if lt-lo < hi-gt {
			quickSort(s, lo, lt-1)
			lo = gt + 1
		} else {
			quickSort(s, gt+1, hi)
			hi = lt - 1
		}
	}
}

// partition partitions s[lo:hi+1] around a median-of-three pivot and
// returns the bounds of the region equal to the pivot: [lt, gt].
func partition[T cmp.Ordered](s []T, lo, hi int) (lt, gt int) {
	mid := lo + (hi-lo)/2
	if s[mid] < s[lo] {
		s[mid], s[lo] = s[lo], s[mid]
	}
	if s[hi] < s[lo] {
		s[hi], s[lo] = s[lo], s[hi]
	}
	if s[hi] < s[mid] {
		s[hi], s[mid] = s[mid], s[hi]
	}
	pivot := s[mid]

	lt, gt = lo, hi
	i := lo
	for i <= gt {
		switch {
		case s[i] < pivot:
			s[lt], s[i] = s[i], s[lt]
			lt++
			i++
		case s[i] > pivot:
			s[i], s[gt] = s[gt], s[i]
			gt--
		default:
			i++
		}
	}
	return lt, gt
}

func insertionSort[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}
