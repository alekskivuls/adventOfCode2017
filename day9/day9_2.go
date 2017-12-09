package main

import (
	"bufio"
	"os"
	"fmt"
)

const (
	NORMAL = iota
	GARBAGE = iota
	IGNORE = iota
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		state := NORMAL
		count, garbageCount, sum := 0, 0, 0
		for _, rune := range scanner.Text() {
			switch state {
			case NORMAL:
				switch rune {
				case '{':
					count++
				case '}':
					sum += count
					count--
				case '<':
					state = GARBAGE
				}
			case GARBAGE:
				switch rune {
				case '!':
					state = IGNORE
				case '>':
					state = NORMAL
				default:
					garbageCount++
				}
			case IGNORE:
				state = GARBAGE
			}
		}
		fmt.Printf("Total sum is: %d\n", sum)
		fmt.Printf("Total size of garbage is: %d\n", garbageCount)
	}
}
