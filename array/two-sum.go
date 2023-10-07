package main

func twoSum(nums []int, target int) []int {
	// var i, j int
	// var result []int

	resultmap := make(map[int]int)

	for i, value := range nums {

		data := target - value

		if index, seen := resultmap[data]; seen {
			return []int{index, i}
		}

		resultmap[value] = i

	}

	// if len(nums) >= 2 && len(nums) <= 10000 {
	// 	for i, j = 0, i+1; i < len(nums)-1; {
	// 		if nums[i]+nums[j] == target {
	// 			result = nil
	// 			result = append(result, i)
	// 			result = append(result, j)
	// 		}
	// 		i++
	// 		j = i + 1
	// 	}
	// }
	// return result

	return nil
}

// func main() {
// 	fmt.Println(twoSum([]int{3, 2, 4, 3}, 6))
// }
