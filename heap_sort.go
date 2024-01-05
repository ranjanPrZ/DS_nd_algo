package main

import (
	"fmt"
)

func heapSort(A []int) {

	// 1. build the heap, get it's size
	buildMaxHeap(A)
	size := len(A)

	for size > 1 {
		// 2. repeatedly swap first and last element, decrement size of considered list
		A[0], A[size-1] = A[size-1], A[0]
		size -= 1

		// 3. siftDown on first element to put it in place
		siftDown(A[:size], 0)
	}
}


// build a Max heap from a given array.
// a max heap adheres to the properties of:
//	1. being filled left to right
//	2. no child node is greater than its parents
func buildMaxHeap(A []int) {
	
	// since all leaf nodes are considered valid heaps, we start from the last
	// non-leaf node, and work right to left in the array
	for i := (len(A) / 2) - 1; i >= 0; i -= 1 {
		siftDown(A, i)
	}
}

// given an index i of A, sift it downwards, should it be smaller than its children
func siftDown(A []int, i int) {
	largest := i
	l, r := getChildIndices(i)
	if l < len(A) && A[l] > A[i] {
		largest = l
	}
	if r < len(A) && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		siftDown(A, largest)
	}
}


func getParentIndex(childIndex int) int {
	if childIndex % 2 == 0 {
		return (childIndex - 2) / 2
	} else {
		return (childIndex - 1) / 2
	}
}


// return child indices as (left-index, right-index)
func getChildIndices(parentIndex int) (int, int) {
	return (2 * parentIndex + 1), (2 * parentIndex + 2)
}

func main() {
	c := []int{13, 48, 1, 4, -4, -3, 20, 0, 100, 99, 11, 3, 7}
	fmt.Println("unsorted:\t\t", c)
	buildMaxHeap(c)
	fmt.Println("buildMaxHeap:\t\t", c)

	d := []int{13, 48, 1, 4, -4, -3, 20, 0, 100, 99, 11, 3, 7}
	heapSort(d)
	fmt.Println("heapSort:\t\t", d)
}
