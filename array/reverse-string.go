package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// func reverse(x int) int {
// 	result := 0
// 	i := len(strconv.Itoa(x)) - 1
// 	for x > 0 {
// 		tmp := x % 10
// 		result = result + tmp*int(math.Pow10(i))
// 		x = x / 10
// 		i--
// 	}
// 	return result
// }

func reverse(x int) int {
	result := big.NewInt(0)
	i := len(strconv.Itoa(x)) - 1
	if big.NewInt(int64(x)).Sign() < 0 {
		// negative number
		for x > 0 {
			tmp := new(big.Int).Mod(big.NewInt(int64(x)), big.NewInt(10))
			result := new(big.Int).Add(
				result, new(big.Int).Mul(tmp, big.NewInt(int64(math.Pow10(i)))))
			x = x / 10
		}
	} else {
		// positive numner
	}

	return int(result.Int64())

	return result
}

func main() {
	fmt.Println(reverse(01234567))
}
