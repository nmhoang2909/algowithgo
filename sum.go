package main

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}

/*
sum({}) = 0
sum({1,2,3}) = 1 + sum({2,3}) = 1 + 2 + sum({3}) = 1 + 2 + 3 + sum({})
*/
