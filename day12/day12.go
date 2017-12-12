package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		file, _ := os.Open(os.Args[1])
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	pipes := make(map[int][]int)
	for scanner.Scan() {
		line := strings.Fields(strings.Replace(scanner.Text(), ",", "", -1))
		key, _ := strconv.Atoi(line[0])
		values := []int{}
		for _, str := range line[2:] {
			val, _ := strconv.Atoi(str)
			values = append(values, val)
		}
		pipes[key] = values
	}

	fmt.Println(countGroups(pipes))
}

func countGroups(pipes map[int][]int) int {
	set := make(map[int]struct{})
	groupCount := 0
	for i := 0; i < len(pipes); i = getNextValNotInSet(set, i) {
		queue := append([]int{}, i)
		for i := 0; i < len(queue); i++ {
			for _, val := range pipes[queue[i]] {
				if _, contains := set[val]; !contains {
					queue = append(queue, val)
					set[val] = struct{}{}
				}
			}
		}
		groupCount++
	}
	return groupCount
}

func getNextValNotInSet(set map[int]struct{}, val int) int {
	for { //Can't put the following in loop constraints since set[val] is evaluated before looping
		if _, contains := set[val]; !contains {
			break
		}
		val++
	}
	return val
}
