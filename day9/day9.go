package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//stack := []rune{}
		count := 0
		garbageCount := 0
		notPrev := false
		inGarbage := false
		total := 0
		for _, rune := range scanner.Text() {
			if !inGarbage {
				switch rune {
				case '<':
					inGarbage = true
				case '{':
					count++
				case '}':
					total += count
					count--
				}
			} else {
				if !notPrev {
					switch rune {
					case '!':
						notPrev = true
					case '>':
						inGarbage = false
					default:
						garbageCount++
					}
				} else {
					notPrev = false
				}
			}
		}
		fmt.Printf("Groups total value is: %d\n", total)
		fmt.Printf("Total garbage count is: %d\n", garbageCount)
	}
}
