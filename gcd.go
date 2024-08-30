package main

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)

	// for {
	// 	if b == 0 {
	// 		return a
	// 	}
	// 	a, b = b, a%b
	// }

	// for b != 0 {
	// 	a, b = b, a%b
	// }
	// return a
}
