package main

import (
	"fmt"
)

func rotateArray(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr // Return an empty array if the input slice is empty
	}

	n := len(arr)
	k %= n // Ensure k is within the range of the slice length

	// Create a new slice to store the rotated elements
	rotated := make([]int, n)

	// Copy the elements that should be at the end after rotation
	copy(rotated, arr[n-k:])

	// Copy the remaining elements to their new positions
	copy(rotated[k:], arr[:n-k])

	return rotated
}

func main() {
	originalArray := []int{1, 2, 3, 4, 5}
	k := 2
	result := rotateArray(originalArray, k)
	fmt.Println(result) // Output: [4 5 1 2 3]
}
