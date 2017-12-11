package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner.Scan()
	dirs := make(map[string]int)
	for _, dir := range strings.Split(scanner.Text(), ",") {
		dirs[dir]++
	}
	reduce(dirs, "se", "nw")
	reduce(dirs, "sw", "ne")
	combine(dirs, "ne", "nw", "n")
	combine(dirs, "se", "sw", "s")
	reduce(dirs, "n", "s")

	sum := 0
	for _, val := range dirs {
		sum += val
	}
	fmt.Println(sum)
}

func reduce(dirs map[string]int, dir string, opp string) {
	if dirs[dir] > dirs[opp] {
		dirs[dir] = dirs[dir] - dirs[opp]
		dirs[opp] = 0
	} else {
		dirs[opp] = dirs[opp] - dirs[dir]
		dirs[dir] = 0
	}
}

func combine(dirs map[string]int, dir1 string, dir2 string, combination string) {
	if dirs[dir1] > dirs[dir2] {
		dirs[combination] += dirs[dir2]
		dirs[dir1] = dirs[dir1] - dirs[dir2]
		dirs[dir2] = 0
	} else {
		dirs[combination] += dirs[dir1]
		dirs[dir2] = dirs[dir2] - dirs[dir1]
		dirs[dir1] = 0
	}
}
