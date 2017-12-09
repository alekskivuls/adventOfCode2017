package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registers := make(map[string]int)
	overallMax := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if evaluateCondition(registers, line[4], line[5], line[6]) {
			val := modifyRegisters(registers, line[0], line[1], line[2])
			if val > overallMax {
				overallMax = val
			}
		}
	}

	max := 0
	for _, val := range registers {
		if val > max {
			max = val
		}
	}
	fmt.Printf("Max right now is %d\n", max)
	fmt.Printf("Overall max was %d\n", overallMax)
}

func evaluateCondition(registers map[string]int, left string, comparator string, right string) bool {
	leftVal := registers[left]
	rightVal, _ := strconv.Atoi(right)
	switch comparator {
	case "==":
		return leftVal == rightVal
	case "!=":
		return leftVal != rightVal
	case "<":
		return leftVal < rightVal
	case "<=":
		return leftVal <= rightVal
	case ">":
		return leftVal > rightVal
	case ">=":
		return leftVal >= rightVal
	default:
		panic("Comparator not found")
	}
}

func modifyRegisters(registers map[string]int, left string, operator string, right string) int {
	rightVal, _ := strconv.Atoi(right)
	switch operator {
	case "inc":
		registers[left] += rightVal
	case "dec":
		registers[left] -= rightVal
	}
	return registers[left]
}
