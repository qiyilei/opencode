// Package quicksort provides an in-place quicksort implementation.
package quicksort

import "cmp"

// Sort sorts the given slice in ascending order using quicksort.
// It works in place and runs in O(n log n) time on average.
func Sort[T cmp.Ordered](s []T) {
	quickSort(s, 0, len(s)-1)
}

func quickSort[T cmp.Ordered](s []T, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(s, lo, hi)
	quickSort(s, lo, p-1)
	quickSort(s, p, hi)
}

// partition splits s[lo:hi+1] around a pivot value and returns the index
// separating the two parts. It uses the middle element as the pivot and a
// Hoare-style two-pointer scan, which degrades gracefully on sorted input.
func partition[T cmp.Ordered](s []T, lo, hi int) int {
	pivot := s[lo+(hi-lo)/2]
	i, j := lo, hi
	for i <= j {
		for s[i] < pivot {
			i++
		}
		for s[j] > pivot {
			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	return i
}
