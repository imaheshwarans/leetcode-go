package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	var result int
	val := x

	for x > 0 {
		result = result*10 + (x % 10)
		x = x / 10
	}

	return result == val
}

func isPalindrome2(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}

func isPalindrome1(x int) bool {
	var reversedNum int

	// If x is a negative number it is not a palindrome
	// If x % 10 = 0, in order for it to be a palindrome the first digit should also be 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x = x / 10
	}

	// If x is equal to reversed number then it is a palindrome
	// If x has odd number of digits, dicard the middle digit before comparing with x
	// Example, if x = 23132, at the end of for loop x = 23 and reversedNum = 231
	// So, reversedNum/10 = 23, which is equal to x
	return x == reversedNum || x == reversedNum/10
}

func main() {

	fmt.Println(isPalindrome(10))

}
