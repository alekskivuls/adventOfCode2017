package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	list := []int{}
	scanner.Scan()
	for _, str := range strings.Fields(scanner.Text()) {
		num, _ := strconv.Atoi(str)
		list = append(list, num)
	}
	count, loopSize := countBalancing(list)
	fmt.Printf("Count: %d\nLoop size: %d\n", count, loopSize)
}

func countBalancing(list []int) (count int, loopSize int) {
	prev, prevVal := 0, 0
	m := make(map[string]int)
	m[fmt.Sprint(list)] = count

	for len(m) > prev {
		prev = len(m)
		maxIndex := 0
		for index, val := range list {
			if val > list[maxIndex] {
				maxIndex = index
			}
		}

		max := list[maxIndex]
		list[maxIndex] = 0
		for i := 1; i <= max; i++ {
			list[(maxIndex + i) % len(list)]++
		}
		count++
		prevVal = m[fmt.Sprint(list)]
		m[fmt.Sprint(list)] = count
	}
	loopSize = count - prevVal
	return
}
