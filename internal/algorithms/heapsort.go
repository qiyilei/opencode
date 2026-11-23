package algorithms

import "cmp"

// HeapSort sorts s in ascending order in place using heap sort.
// It runs in O(n log n) time and O(1) extra space.
func HeapSort[T cmp.Ordered](s []T) {
	n := len(s)

	// Build a max-heap from the bottom up.
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(s, i, n)
	}

	// Repeatedly move the max element to the end of the heap.
	for end := n - 1; end > 0; end-- {
		s[0], s[end] = s[end], s[0]
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
