package main

import (
	"math"
)

/*
1110
res = res + 0 * 2^0 = 0
res = res + 1 * 2^1 = 2
res = res + 1 * 2^2 = 6
res = res + 1 * 2^3 = 14
*/
func BaseToDec(value string, base int) int {
	charset := "0123456789ABCDEF"
	res := 0
	power := 0
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		res += index * int(math.Pow(float64(base), float64(power)))
		power++
	}
	return res
}
