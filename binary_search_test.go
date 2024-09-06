package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			// arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target: 10,
			want:   9,
		},
		{
			// arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target: 15,
			want:   14,
		},
		{
			// arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target: 1,
			want:   0,
		},
		{
			// arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target: 1000000,
			want:   999999,
		},
		{
			target: 1000001,
			want:   -1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.target), func(t *testing.T) {
			arrLen := 1000000
			arr := make([]int, arrLen)
			for i := 0; i < arrLen; i++ {
				arr[i] = i + 1
			}
			got := BinarySearch(arr, test.target)
			assert.Equal(t, test.want, got)
		})
	}
}

func BinarySearch(arr []int, target int) int {
	l, h := 0, len(arr)-1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] < target {
			l = mid + 1
		}
		if arr[mid] > target {
			h = mid - 1
		}
		if arr[mid] == target {
			return mid
		}
	}
	return -1
}
