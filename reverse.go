package main

func Reverse(word string) string {
	var ret string
	for _, r := range word {
		ret = string(r) + ret
	}
	return ret
}
