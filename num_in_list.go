package main

func NumInList(list []int, num int) bool {
	for _, item := range list {
		if item == num {
			return true
		}
	}
	return false
}
