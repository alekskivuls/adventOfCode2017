package main

import "testing"

var TEST_MIN_MAX_SPREADSHEET = [][]int{
	{5,1,9,5},
	{7,5,3},
	{2,4,6,8},
	}

var TEST_EVENLY_DIVISIBLE_SPREADSHEET = [][]int{
	{5,9,2,8},
	{9,4,7,3},
	{3,8,6,5},
}

func TestCheckSumRowMinMax(t *testing.T) {
	sum := 0
	for _, row := range TEST_MIN_MAX_SPREADSHEET {
		sum += CheckSumRowMinMax(row)
	}

	expected := 18
	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}

func TestCheckSumRowEvenlyDivisible(t *testing.T) {
	sum := 0
	for _, row := range TEST_EVENLY_DIVISIBLE_SPREADSHEET {
		sum += CheckSumRowEvenlyDivisible(row)
	}

	expected := 9
	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}