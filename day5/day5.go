package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	list := []int{}
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		list = append(list, num)
	}
	listCopy := make([]int, len(list))
	copy(listCopy, list)

	fmt.Printf("Number of regular steps: %d\n", countSteps(list, false))
	fmt.Printf("Number of steps with decrement: %d\n", countSteps(listCopy, true))
}

func countSteps(list []int, dec bool) int {
	var count, index, jump int
	for index + jump >= 0 && index + jump < len(list) {
		index += jump
		jump = list[index]
		if dec {
			if jump >= 3 {
				list[index]--
			} else {
				list[index]++
			}
		} else {
			list[index]++
		}
		count++
	}
	return count
}