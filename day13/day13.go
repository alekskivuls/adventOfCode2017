package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanners := make(map[int]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		num, _ := strconv.Atoi(line[0])
		size, _ := strconv.Atoi(line[1])
		scanners[num] = size
	}
	fmt.Printf("Total complexity: %d\n", getTotalComplexity(scanners))
	fmt.Printf("Safe delay: %d\n", getDelay(scanners))
}

func getTotalComplexity(scanners map[int]int) int {
	sum := 0
	for id, size := range scanners {
		sum += getComplexity(id, size)
	}
	return sum
}

func getDelay(scanners map[int]int) int {
	delay := 0
	safe := false
	for !safe {
		safe = true
		for id, size := range scanners {
			if getComplexity(id+delay, size) != 0 {
				safe = false
				delay++
			}
		}
	}
	return delay
}

func getComplexity(id int, size int) int {
	sum := 0
	if id%(2*size-2) == 0 {
		sum += id * size
	}
	return sum
}
