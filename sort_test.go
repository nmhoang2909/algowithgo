package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{87, 23, 56, 12, 98, 45, 78, 34, 67, 10},
			want:  []int{10, 12, 23, 34, 45, 56, 67, 78, 87, 98},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			got := SelectionSort(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		isSwap := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return nums
}

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		j := i
		for j > 0 && nums[j-1] > nums[j] {
			nums[j-1], nums[j] = nums[j], nums[j-1]
			j--
		}
	}
	return nums
	// for i := 1; i < len(nums); i++ {
	// 	temp := nums[i]
	// 	j := i - 1
	// 	for j >= 0 && nums[j] > temp {
	// 		nums[j], nums[j+1] = nums[j+1], nums[j]
	// 		j--
	// 	}
	// 	nums[j+1] = temp
	// 	fmt.Printf("nums: %v\n", nums)
	// }
	// return nums
}

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
		fmt.Printf("nums: %v\n", nums)
	}
	return nums
}
