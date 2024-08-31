package main

import (
	"fmt"
)

func FizzBuzz(number int) {
	for i := 1; i <= number; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Print(i)
		}
		if i != number {
			fmt.Print(", ")
		}
	}
	fmt.Println("")
}
