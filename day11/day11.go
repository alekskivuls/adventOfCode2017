package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coordinate struct {
	X int
	Y int
	Z int
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		file, _ := os.Open(os.Args[1])
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Scan()
	c := coordinate{}
	max := 0
	for _, dir := range strings.Split(scanner.Text(), ",") {
		switch dir {
		case "n":
			c.Y++
			c.Z--
		case "s":
			c.Z++
			c.Y--
		case "ne":
			c.X++
			c.Z--
		case "sw":
			c.Z++
			c.X--
		case "nw":
			c.Y++
			c.X--
		case "se":
			c.X++
			c.Y--
		}
		if Dist(c) > max {
			max = Dist(c)
		}
	}

	fmt.Printf("Current distance is: %d\n", Dist(c))
	fmt.Printf("Max distance was: %d\n", max)
}

func Dist(c coordinate) int {
	return (abs(c.X) + abs(c.Y) + abs(c.Z)) / 2
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}
