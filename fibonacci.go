package main

var couter = make(map[int]int)

// value 0 1 1 2 3 5 8 13
// index 0 1 2 3 4 5 6 7
func Fibonacci(num int) int {
	couter[num]++
	if num <= 1 {
		return num
	}

	// return Fibonacci(num-1) + Fibonacci(num-2)

	/*
	  curNum = 2
	  nMinus2 = curNum - 2 = 0
	  nMinus1 = curNum - 1 = 1
	  curValue = nMinus2 + nMinus1 = 5
	  nMinus2 = nMinus1 = 3
	  nMinus1 = curValue = 5
	  curNum++ = 5
	*/

	var res int
	nMinus2 := 0
	nMinus1 := 1
	for i := 2; i <= num; i++ {
		res = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = res
	}

	return res
}
