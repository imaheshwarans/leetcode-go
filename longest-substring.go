package main

// import "fmt"

// func lengthOfLongestSubstring(s string) int {
// 	var i int
// 	r := []rune(s)
// 	var substring []string

// 	for i = 0; i < len(s); i++ {
// 		j := i + 1
// 		substr := string(r[i])
// 		for ; j < len(s)-i; j++ {
// 			if r[i] != r[j] {
// 				substr = substr + string(r[j])
// 			} else {
// 				break
// 			}
// 		}
// 		substring = append(substring, substr)
// 	}
// 	max := 0
// 	for _, str := range substring {
// 		if max < len(str) {
// 			max = len(str)
// 		}
// 	}
// 	return max
// }

// func main() {

// 	str := "pwwkew"
// 	// var substring []string
// 	var max, i int
// 	tmpMap := make(map[rune]bool)

// 	r := []rune(str)
// 	tmp := r[i]
// 	max++
// 	tmpmax := max
// 	for i = 1; i < len(str); i++ {
// 		if tmp != r[i] {
// 			if seen := tmpMap[tmp]; seen {
// 				tmpMap = nil
// 				tmpmax = 0
// 				break
// 			} else {
// 				tmpMap[tmp] = true
// 				tmp = r[tmpmax]
// 				tmpmax++
// 			}
// 		}

// 		if max < tmpmax {
// 			max = tmpmax - 1
// 		}
// 	}

// 	fmt.Println(max)

// }
