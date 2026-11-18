// Package quicksort provides an in-place generic quicksort implementation.
package quicksort

import (
	"cmp"
	"math/rand/v2"
)

// Sort sorts arr in place in ascending order using randomized-pivot quicksort.
func Sort[T cmp.Ordered](arr []T) {
	sortRange(arr, 0, len(arr)-1)
}

func sortRange[T cmp.Ordered](arr []T, left, right int) {
	if left >= right {
		return
	}

	// Random pivot keeps already sorted input from degrading to O(n^2).
	pivotIdx := left + rand.IntN(right-left+1)
	arr[left], arr[pivotIdx] = arr[pivotIdx], arr[left]
	pivot := arr[left]

	i, j := left+1, right
	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Put the pivot into its final position.
	arr[left], arr[j] = arr[j], arr[left]

	sortRange(arr, left, j-1)
	sortRange(arr, j+1, right)
}
