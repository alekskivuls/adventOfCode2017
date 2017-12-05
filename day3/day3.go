package main

import "fmt"

func main() {
	fmt.Println(GetRingDist(265149))
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