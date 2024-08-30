package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var gcd int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		gcd = GCD(a, b)
		fmt.Println(gcd)
	}
}
