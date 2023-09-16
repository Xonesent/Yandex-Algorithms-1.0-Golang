package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	tdn := strings.Fields(params[0])
	t, _ := strconv.Atoi(tdn[0])
	d, _ := strconv.Atoi(tdn[1])
	n, _ := strconv.Atoi(tdn[2])
	start := [4]int{0, 0, 0, 0}

	for i := 1; i <= n; i++ {
		start = zone(start, t)
		xy := strings.Fields(params[i])
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		finish := zone([4]int{x + y, x + y, x - y, x - y}, d)
		start = anspoint(start, finish)
	}

	var anslist [][]int

	for side1 := start[0]; side1 <= start[1]; side1++ {
		for side3 := start[2]; side3 <= start[3]; side3++ {
			if (side1+side3)%2 == 0 {
				x := (side1 + side3) / 2
				y := side1 - x
				anslist = append(anslist, []int{x, y})
			}
		}
	}

	fmt.Println(len(anslist))
	for _, coord := range anslist {
		fmt.Println(coord[0], coord[1])
	}
}

func zone(a [4]int, t int) [4]int {
	return [4]int{a[0] - t, a[1] + t, a[2] - t, a[3] + t}
}

func anspoint(a [4]int, b [4]int) [4]int {
	return [4]int{max(a[0], b[0]), min(a[1], b[1]), max(a[2], b[2]), min(a[3], b[3])}
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
