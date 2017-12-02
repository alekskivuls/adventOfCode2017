package day2

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0
	checksum2 := 0
	for scanner.Scan() {
		var nums []int
		for _, str := range strings.Fields(scanner.Text()) {
			num,_ := strconv.Atoi(string(str))
			nums = append(nums, num)
		}
		checksum += CheckSumRowMinMax(nums)
		checksum2 += CheckSumRowEvenlyDivisible(nums)
	}
	fmt.Printf("Part one: %d\n", checksum)
	fmt.Printf("Part two: %d\n", checksum2)

}

func CheckSumRowMinMax(nums []int) int {
	min := nums[0]
	max := min
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max - min
}

func CheckSumRowEvenlyDivisible(nums []int) int {
	sort.Ints(nums)
	for index, denominator := range nums {
		for _, numerator := range nums[index+1:] {
			if numerator % denominator == 0 {
				return numerator / denominator
			}
		}
	}
	panic("No valid checksum for row found")
}
