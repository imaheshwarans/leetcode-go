package main

import (
	"fmt"
	"sort"
)

func binarySearchTree(array []int, search int) bool {

	if !sort.IntsAreSorted(array) {
		sort.Ints(array)
	}

	left, right := 0, len(array)-1

	// found := false
	for left <= right {
		middle := (left + right) / 2
		if array[middle] == search {
			return true
		}
		if array[middle] < middle {
			left = middle + 1
		} else {
			right = middle - 1
		}

	}

	return false
}

// func binarySearch(arr []int, target int) int {
// 	left, right := 0, len(arr)-1

// 	for left <= right {
// 		mid := (left + right) / 2

// 		if arr[mid] == target {
// 			return mid
// 		}

// 		if arr[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return -1
// }

func main() {

	var find []int = []int{2, 3, 4, 5, 6, 7, 8}

	fmt.Println(binarySearchTree(find, 1235))

}
