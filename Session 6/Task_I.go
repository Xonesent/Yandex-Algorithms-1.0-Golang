package main

import (
	"fmt"
	"sort"
)

var class []int
var n, r, c int

func main() {
	fmt.Scan(&n, &r, &c)
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		class = append(class, value)
	}
	sort.Ints(class)

	l := 0
	ri := class[n-1] - class[0]

	for l < ri {
		m := (l + ri) / 2
		if check(m) {
			l = m + 1
		} else {
			ri = m
		}
	}
	fmt.Println(l)
}
func check(m int) bool {
	var left int
	var right int = n - 1
	var count int
	for left <= right-c+1 {
		if class[left+c-1]-class[left] <= m {
			count++
			left += c
		} else {
			left++
		}
	}
	return count < r
}
