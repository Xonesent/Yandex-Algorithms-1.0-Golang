package main

import (
	"fmt"
)

var list [][]int
var n, l int

func main() {
	fmt.Scan(&n, &l)

	for i := 1; i <= n; i++ {
		var x1, d1, a, c, m int
		fmt.Scan(&x1, &d1, &a, &c, &m)
		list = append(list, []int{x1})
		for idx := 1; idx < l; idx++ {
			x1 = x1 + d1
			list[i-1] = append(list[i-1], x1)
			d1 = (a*d1 + c) % m
		}
	}
	for i := 0; i < n-1; i++ {
		for idx := i + 1; idx < n; idx++ {
			fmt.Println(poisk(list[i], list[idx]))
		}
	}
}

func poisk(list1 []int, list2 []int) int {
	var left int = min(list1[0], list2[0])
	var right int = max(list1[l-1], list2[l-1])
	for left < right {
		m := (left + right + 1) / 2
		count1 := points(list1, m)
		count2 := points(list2, m)
		count := count1 + count2
		if count >= l {
			right = m - 1
		} else {
			left = m
		}
	}
	return left
}

func points(list []int, m int) int {
	var left1 int = 0
	var right1 int = len(list)
	for left1 < right1 {
		mid1 := (left1 + right1) / 2
		if list[mid1] < m {
			left1 = mid1 + 1
		} else {
			right1 = mid1
		}
	}
	return left1
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
