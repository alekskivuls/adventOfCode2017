package main

import "testing"

func TestSumIfSameNum(t *testing.T) {
	var nums []int
	nums = []int{1,1,2,2}
	expect(t, SumIfSameNum(nums, 1), 3)
	nums = []int{1,1,1,1}
	expect(t, SumIfSameNum(nums, 1), 4)
	nums = []int{1,2,3,4}
	expect(t, SumIfSameNum(nums, 1), 0)
	nums = []int{9,1,2,1,2,1,2,9}
	expect(t, SumIfSameNum(nums, 1), 9)

	nums = []int{1,2,1,2}
	expect(t, SumIfSameNum(nums, len(nums)/2), 6)
	nums = []int{1,2,2,1}
	expect(t, SumIfSameNum(nums, len(nums)/2), 0)
	nums = []int{1,2,3,4,2,5}
	expect(t, SumIfSameNum(nums, len(nums)/2), 4)
	nums = []int{1,2,3,1,2,3}
	expect(t, SumIfSameNum(nums, len(nums)/2), 12)
	nums = []int{1,2,1,3,1,4,1,5}
	expect(t, SumIfSameNum(nums, len(nums)/2), 4)
}

func expect(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", actual, expected)
	}
}