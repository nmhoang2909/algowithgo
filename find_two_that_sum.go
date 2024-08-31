package main

import "fmt"

func FindTwoThatSum(numbers []int, sum int) (i int, j int) {
	for i, v1 := range numbers {
		for j := len(numbers) - 1; j >= 0; j-- {
			fmt.Printf("j: %v\n", j)
			if i == j {
				continue
			}
			if v1+numbers[j] != sum {
				continue
			}
			return i, j
		}

	}
	return -1, -1
}
