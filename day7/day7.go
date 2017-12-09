package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"sort"
)

type node struct {
	val int
	parent string
	children []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nodes := make(map[string]*node)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		name := line[0]
		val,_ := strconv.Atoi(line[1][1:len(line[1])-1])
		var children []string
		if len(line) > 2 {
			for _, str := range line[3:] {
				children = append(children, strings.Replace(str, ",", "", -1))
			}
		}
		nodes[name] = &node{val, "", children}
	}
	for name, val := range nodes {
		for _, child := range val.children {
			nodes[child].parent = name
			if len(val.children) == 1 {
				panic("test")
			}
		}
	}

	root := findRoot(nodes)
	fmt.Printf("%s: %+v\n", root, *nodes[root])
	fmt.Printf("%d\n", -findOffBalanceWeight(nodes, root))
}

func findRoot(nodes map[string]*node) string {
	for k, v := range nodes {
		if v.parent == "" {
			return k
		}
	}
	return ""
}

func findOffBalanceWeight(nodes map[string]*node, node string) (int) {
	if nodes[node].children != nil {
		childVals := []int{}
		prevNodeWeight := make(map[int]int)
		for _, child := range nodes[node].children {
			weight := findOffBalanceWeight(nodes, child)
			prevNodeWeight[weight] = nodes[child].val
			childVals = append(childVals, weight)
		}
		sort.Ints(childVals)

		low := childVals[0]
		mid := childVals[1]
		high := childVals[len(childVals)-1]

		if low < 0 {
			return low
		} else {
			if low != high {
				if low != mid {
					prev := prevNodeWeight[low]
					return -(mid - low + prev)
				} else {
					prev := prevNodeWeight[high]
					return -(low - high + prev)
				}
			} else {
				return sumInts(childVals) + nodes[node].val
			}
		}
	} else {
		return nodes[node].val
	}
}

func sumInts(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}