package main

import "fmt"

func BaseToBase(value string, base, newBase int) string {
	dec := BaseToDec(value, base)
	fmt.Printf("dec: %v\n", dec)
	return DecToBase(dec, newBase)
}
