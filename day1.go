package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nums []int
	for _, r := range scanner.Text() {
		num,_ := strconv.Atoi(string(r))
		nums = append(nums, num)
	}
	fmt.Printf("Part one: %d\n", SumIfSameNum(nums, 1))
	fmt.Printf("Part two: %d\n", SumIfSameNum(nums, len(nums)/2))
}

func SumIfSameNum(nums []int, offset int) int {
	sum := 0
	for index, num := range nums {
		if nums[(index + offset) % len(nums)] == num {
			sum += num
		}
	}
	return sum
}
