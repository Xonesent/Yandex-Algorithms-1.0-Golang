package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var n, m, l int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ = strconv.Atoi(params[0])
	m, _ = strconv.Atoi(params[1])
	l, _ = strconv.Atoi(params[2])

	left := 0
	right := min(n, m) / 2

	for left < right {
		mid := (left + right + 1) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Println(left)
}

func check(mid int) bool {
	return 2*max(n, m)*mid+2*(min(n, m)-2*mid)*mid <= l
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
