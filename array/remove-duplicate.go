package main

import "fmt"

func removeDuplicates(nums []int) int {
	// length := 0
	// unique := make(map[int]bool)
	// for i := 0; i < len(nums); i++ {
	// 	if seen := unique[nums[i]]; seen {
	// 		// remove from that index
	// 		// and left shift ?
	// 		for j := i; j < len(nums)-1; j++ {
	// 			nums[j] = nums[j+1]
	// 		}
	// 	}
	// 	unique[nums[i]] = true
	// 	if nums[i] >= 0 {
	// 		length++
	// 	}
	// }
	j := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i+1]
			j++
		}
	}

	fmt.Println(nums)

	return j
}

func main() {

	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

}
