package main

func Factor(primes []int, num int) []int {
	var res []int
	for _, prime := range primes {
		for num%prime == 0 {
			res = append(res, prime)
			num = num / prime
		}

	}
	if num > 1 {
		res = append(res, num)
	}
	return res
}
