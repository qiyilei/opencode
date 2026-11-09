// Package heapsort provides an in-place heapsort implementation.
package heapsort

import "cmp"

// Sort sorts the given slice in ascending order using heapsort.
// It works in place and runs in O(n log n) time in the worst case.
func Sort[T cmp.Ordered](s []T) {
	n := len(s)

	// Build a max-heap: start from the last non-leaf node and sift down.
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(s, i, n)
	}

	// Move the maximum to the end of the unsorted region, then restore the heap.
	for end := n - 1; end > 0; end-- {
		s[end], s[0] = s[0], s[end]
		siftDown(s, 0, end)
	}
}

// siftDown restores the max-heap property for the subtree rooted at root,
// considering only the first n elements of s.
func siftDown[T cmp.Ordered](s []T, root, n int) {
	for {
		child := 2*root + 1
		if child >= n {
			return
		}
		// Pick the larger of the two children.
		if child+1 < n && s[child+1] > s[child] {
			child++
		}
		if s[root] >= s[child] {
			return
		}
		s[root], s[child] = s[child], s[root]
		root = child
	}
}
