package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumInList(t *testing.T) {
	tests := []struct {
		list []int
		num  int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2, 3, 4, -1}, -1, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 8, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 2, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 5, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 4, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 1, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 8, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 9, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 3, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 0, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 88, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 23, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 44, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 123, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 321, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 32, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 7, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 6, false},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, -1, true},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, 1, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.list), func(t *testing.T) {
			actual := NumInList(test.list, test.num)
			assert.Equal(t, test.want, actual)
		})
	}
}
