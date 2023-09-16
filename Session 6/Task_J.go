package main

import (
	"fmt"
)

var list [][]int
var n, l int

func main() {
	fmt.Scan(&n, &l)
	for i := 1; i <= n; i++ {
		list = append(list, []int{})
		for idx := 0; idx < l; idx++ {
			var value int
			fmt.Scan(&value)
			list[i-1] = append(list[i-1], value)
		}
	}
	for i := 0; i < n-1; i++ {
		for idx := i + 1; idx < n; idx++ {
			fmt.Println(poisk(list[i], list[idx]))
		}
	}
}
func poisk(list1 []int, list2 []int) int {
	var left1, left2 int
	var ans int
	for left1+left2 < l {
		if list1[left1] <= list2[left2] {
			ans = list1[left1]
			left1++
		} else {
			ans = list2[left2]
			left2++
		}
	}
	return ans
}
