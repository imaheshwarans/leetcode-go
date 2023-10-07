// Golang program to reverse a string
package main

import "fmt"

// importing fmt

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func MissingNo(a []int, n int) int {
	var i, total int
	total = (n + 1) * (n + 2) / 2
	for i = 0; i < n; i++ {
		total -= a[i]
	}

	return total
}

func missingNumber(array []int, n int) int {
	n++
	total := (n * (n + 1)) / 2
	var sum int = 0
	for _, num := range array {
		sum = sum + num
	}

	fmt.Println("sum ", sum)

	return total - sum
}

func main() {

	array := []int{5, 6, 8, 9}

	fmt.Println(missingNumber(array, len(array)))
}
