package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var n, x, y int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	params := string(file)

	nxy := strings.Fields(string(params))
	n, _ = strconv.Atoi(nxy[0])
	x, _ = strconv.Atoi(nxy[1])
	y, _ = strconv.Atoi(nxy[2])

	left := min(x, y)
	right := min(x, y) * n

	for left < right {
		m := (left + right) / 2
		if check(m) {
			right = m
		} else {
			left = m + 1
		}
	}

	fmt.Println(left)
}

func check(m int) bool {
	return 1+(m-min(x, y))/x+(m-min(x, y))/y >= n
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
