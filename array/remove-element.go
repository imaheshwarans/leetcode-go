package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0

	for _, num := range nums {
		if num != val {
			nums[index] = num
			index += 1
		}
	}
	return index
}

func main() {

	fmt.Println(removeElement([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 4))

}
