package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	anagramCount := 0
	for scanner.Scan() {
		if isValidPass(scanner.Text(), false) {
			count++
		}
		if isValidPass(scanner.Text(), true) {
			anagramCount++
		}
	}
	fmt.Printf("Unique word passphrase count: %d\n",count)
	fmt.Printf("Unique anagram word passphrase count: %d\n", anagramCount)
}

func isValidPass(str string, anagram bool) bool {
	m := make(map[string]int)
	for _, val := range strings.Fields(str) {
		if anagram {
			s := strings.Split(val, "")
			sort.Strings(s)
			val = strings.Join(s, "")
		}
		if _, ok := m[val]; ok {
			return false
		} else {
			m[val] = 0
		}
	}
	return true
}