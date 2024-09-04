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
			got := MergeSort(test.input)
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

func MergeSort(nums []int) []int {
	le := len(nums)
	if le <= 1 {
		return nums
	}

	mid := le / 2
	leftArr := make([]int, mid)
	rightArr := make([]int, le-mid)

	var j int
	for i := 0; i < le; i++ {
		if i < mid {
			leftArr[i] = nums[i]
		} else {
			rightArr[j] = nums[i]
			j++
		}
	}

	leftArr = MergeSort(leftArr)
	rightArr = MergeSort(rightArr)
	merge(leftArr, rightArr, nums)
	fmt.Printf("nums: %v\n", nums)
	return nums
}

func merge(left, right, nums []int) {
	leftSize := len(left)
	rightSize := len(nums) - leftSize

	var l, r, i int
	for l < leftSize && r < rightSize {
		if left[l] < right[r] {
			nums[i] = left[l]
			i++
			l++
		} else {
			nums[i] = right[r]
			i++
			r++
		}
	}
	for l < leftSize {
		nums[i] = left[l]
		i++
		l++
	}
	for r < rightSize {
		nums[i] = right[r]
		i++
		r++
	}
}
