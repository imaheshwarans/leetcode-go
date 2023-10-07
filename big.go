package main

// import (
// 	"fmt"
// 	"math/big"
// )

// func main() {

// 	// var tmp *big.Int
// 	tmp, _ := new(big.Int).SetString("9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", 10)
// 	var i int

// 	fmt.Println(tmp)

// 	fmt.Println("\n\n")
// 	fmt.Println(
// 		new(big.Int).Add(tmp, big.NewInt(int64(1))))

// 	fmt.Println("\n\n")

// 	// for i < 20 {
// 	// 	tmp = tmp.Add(tmp, big.NewInt(int64(
// 	// 		math.Pow(10, float64(i)))))
// 	// 	i++
// 	// }
// 	curr := 2
// 	for i < 32 {
// 		tmp = new(big.Int).Add(tmp, new(big.Int).Mul(big.NewInt(int64(curr)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(i)), nil)))
// 		i++
// 	}

// 	fmt.Println(tmp)
// 	// 1111111111111111111
// }

// package main

// import (
// 	"fmt"
// 	"math/big"
// )

// func main() {
// 	strNum := "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"

// 	bigNum, success := new(big.Int).SetString(strNum, 10)

// 	// r := []rune(strNum)

// 	fmt.Println(len(bigNum.String()))

// 	for _, value := range bigNum.String() {
// 		fmt.Println(value - '0')
// 	}

// 	if !success {
// 		fmt.Println("Failed to convert the string to a big integer")
// 		return
// 	}

// 	one := big.NewInt(1)

// 	result := new(big.Int).Add(bigNum, one)

// 	fmt.Println("Result:", result)
// }
