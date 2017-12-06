package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetRingDist(265149))
	fmt.Println(GenerateRing(265149))
}

func GetRing(num int) (ring int, val int) {
	for num > getRingValue(ring) {
		ring++
	}
	return ring, getRingValue(ring)
}

//The ring value can be found by (2 * i + 1)^2
func getRingValue(ring int) int {
	return (2 * ring + 1) * (2 * ring + 1)
}

func GetRingDistOffset(num int) int {
	ring, val := GetRing(num)
	offset := ring
	inc := false
	for i := 0; val - i != num; i++ {
		if offset == 0 {
			inc = true
		} else if offset == ring {
			inc = false
		}

		if inc {
			offset++
		} else {
			offset--
		}
	}
	return offset
}

func GetRingDist(num int) int {
	ring, _ := GetRing(num)
	return ring + GetRingDistOffset(num)
}

//Part 2
func GenerateRing(num int) int {
	size := 1000
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, size)
	}

	row, col := size/2, size/2
	arr[row][col] = 1

	for ringSize := 2;; ringSize += 2{
		col++
		row++
		for dir := 0; dir < 4; dir++ {
			for steps := 0; steps < ringSize; steps++ {
				row, col = goForward(dir, row, col)
				arr[row][col] = sumAdjacent(arr, row, col)
				if arr[row][col] > num {
					return arr[row][col]
				}
			}
		}
	}
	return -1
}

func goForward(dir int, row int, col int) (int, int) {
	switch dir {
	case 0:
		return row - 1, col
	case 1:
		return row, col - 1
	case 2:
		return row + 1, col
	case 3:
		return row, col + 1
	default:
		return row, col
	}
}

func sumAdjacent(arr [][]int, x int, y int) (sum int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += arr[x+i][y+j]
		}
	}
	return
}