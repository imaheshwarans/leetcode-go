package main

import (
	"fmt"
)

type Hello struct {
	Name string
}

func sayHelloWorld() {
	fmt.Println("saying hi from sayHelloWorld function")
}

func sayHelloWithName(hello Hello) {
	fmt.Println("Hi ", hello.Name)
}

func (h *Hello) hello() {
	fmt.Println("Hi ", h.Name)
}

func main() {
	// fmt.Println("Hello World!")

	// sayHelloWorld()

	// hello := Hello{
	// 	Name: "Maheshwaran S.",
	// }

	// sayHelloWithName(hello)

	// hello.hello()

	// // var name string

	// name := "hellp"

	// fmt.Println("name ", name)

	// go func() {
	// 	fmt.Println("hi")
	// }()

	// fmt.Println("The time is", time.Now())

	sliceAndArray()

	// printMatrixInSnake()
	// printSpiralOrder()
}

func sliceAndArray() {
	sliceint := [4]int{0, 1, 2, 3}

	fmt.Println(sliceint)

	//sliceint = append(sliceint, 4)

	var slice1 []int

	slice2 := sliceint[1:]
	fmt.Println(slice2)

	slice1 = sliceint[1:3]
	fmt.Println(slice1)

	slice1 = append(slice1, 4)

	fmt.Println(slice1)

	var names []string = []string{"m", "a", "h", "e", "s", "h"}

	fmt.Println(names[1:])
	fmt.Println(names[0:])
	fmt.Println(names[:6])
	fmt.Println(names[2:4])
}

func printMatrixInSnake() {
	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println(matrix)

	i, j := 0, 0

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("\n")
	}

	// snake
	for i = 2; i >= 0; i-- {
		if i%2 == 0 {
			for j = 2; j >= 0; j-- {
				fmt.Print(matrix[i][j])
			}
		} else {
			for j = 0; j < 3; j++ {
				fmt.Print(matrix[i][j])
			}

		}

		// fmt.Println("\n")
	}
}

func printMatrixInSpiral() {
	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println(matrix)

	i, j, n := 0, 0, 3

	// output 123698745
	rowStart := 0
	rowEnd := 2

	columnStart := 0
	columnEnd := 2

	for i = rowStart; i < n; i++ {

		for j = columnStart; j < n; j++ {
			fmt.Print(matrix[i][j])
		}

		k := 0
		if j == columnEnd {
			rowStart--
			for k = i + 1; k < rowEnd; k++ {
				fmt.Print(matrix[k][columnEnd])
			}
		}

		if k == columnEnd {
			k = 0
			columnEnd--
			for k = columnEnd + 1; k > rowStart; k-- {
				fmt.Print(matrix[rowEnd][k])
			}
		}

		if k == rowStart {
			k = 0
			rowEnd--
			for k = columnEnd + 1; k > rowStart; k-- {
				fmt.Print(matrix[rowEnd][k])
			}
		}

		fmt.Println("\n")
	}

	// snake
	for i = 2; i >= 0; i-- {
		if i%2 == 0 {
			for j = 2; j >= 0; j-- {
				fmt.Print(matrix[i][j])
			}
		} else {
			for j = 0; j < 3; j++ {
				fmt.Print(matrix[i][j])
			}

		}

		// fmt.Println("\n")
	}
}

func printSpiralOrder() {
	m, n := 3, 3
	rowStart := 0
	rowEnd := m - 1
	colStart := 0
	colEnd := n - 1

	var X [3][3]int

	X[0] = [3]int{1, 2, 3}
	X[1] = [3]int{4, 5, 6}
	X[2] = [3]int{7, 8, 9}

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i = i + 1 {
			print(X[rowStart][i])
		}

		rowStart = rowStart + 1

		for i := rowStart; i <= rowEnd; i = i + 1 {
			print(X[i][colEnd])
		}

		colEnd = colEnd - 1

		if rowStart <= rowEnd {
			for i := colEnd; i >= colStart; i = i - 1 {
				print(X[rowEnd][i])
			}
			rowEnd = rowEnd - 1
		}

		if colStart <= colEnd {
			for i := rowEnd; i >= rowStart; i = i - 1 {
				print(X[i][colStart])
			}
			colStart = colStart + 1
		}
	}
}
