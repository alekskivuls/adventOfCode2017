package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"encoding/hex"
)

const NUM_ELEMENTS = 256
const NUM_ROUNDS = 64
const BLOCK_SIZE = 16

func main() {
	ring := []int{}
	for i := 0; i < NUM_ELEMENTS; i++ {
		ring = append(ring, i)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := []int{}
	inputTerminator := []int{17, 31, 73, 47, 23}
	for _, rune := range scanner.Text() {
		input = append(input, int(rune))
	}
	input = append(input, inputTerminator...)

	currPos := 0
	skipSize := 0
	for i := 0; i < NUM_ROUNDS; i++ {
		for _, num := range input {
			if num > 1 {
				reverse(ring, currPos, (currPos+num-1)%len(ring))
			}
			currPos = (currPos + skipSize + num) % len(ring)
			skipSize++
		}
	}

	hashRing := []byte{}
	for block := 0; block < NUM_ELEMENTS/BLOCK_SIZE; block++ {
		xorVal := ring[block*BLOCK_SIZE]
		for i := 1; i < BLOCK_SIZE; i++ {
			xorVal = xorVal ^ ring[block*BLOCK_SIZE+i]
		}
		hashRing = append(hashRing, byte(xorVal))
	}
	fmt.Println(hex.EncodeToString(hashRing))
}

func partOne(ring []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	currPos := 0
	for skipSize, str := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(str)
		if num > 1 {
			reverse(ring, currPos, (currPos+num-1)%len(ring))
		}
		currPos = (currPos + skipSize + num) % len(ring)
	}
	fmt.Printf("First two multiplied: %d\n", ring[0]*ring[1])
}

func reverse(slice []int, start int, end int) {
	var dist int
	if end > start {
		dist = end - start
	} else if end < start {
		dist = len(slice) - start + end
	} else {
		return
	}

	for i := 0; i <= dist/2; i++ {
		left := (start + i) % len(slice)
		right := (start + dist - i) % len(slice)
		slice[left], slice[right] = slice[right], slice[left]
	}
}
