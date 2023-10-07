// Golang program to reverse a string
package main

// importing fmt
import "fmt"

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

func reverseFor(str string) (result string) {
	for i, v := range str {
		fmt.Println(i)
		fmt.Println(string(v))
		result = string(v) + result
		fmt.Println(result)
	}
	return
}

func main() {

	// Reversing the string.
	str := "Geeks"

	// returns the reversed string.
	strRev := reverseFor(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
