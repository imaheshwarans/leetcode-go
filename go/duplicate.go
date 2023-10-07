package main

import "fmt"

func findDuplicates(nums []int) []int {
	duplicates := []int{}
	seen := make(map[int]bool)

	fmt.Println(seen)

	for _, num := range nums {
		if seen[num] {
			fmt.Println(seen[num])
			duplicates = append(duplicates, num)
		} else {
			seen[num] = true
		}
	}

	return duplicates
}

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 1, 5}
	duplicates := findDuplicates(arr)

	fmt.Println("Duplicates:", duplicates)
}
