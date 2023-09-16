package main

import "fmt"

var numlist []int
var n, k int
var count int

func main() {
	fmt.Scan(&n, &k)
	var max int
	var maxk int
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		maxk += value
		if value > max {
			max = value
		}
		numlist = append(numlist, value)
	}
	left := 1
	right := max

	if maxk < k {
		fmt.Println(0)
	} else {
		for left < right {
			m := (left + right + 1) / 2
			count = 0
			if check(m) {
				left = m
			} else {
				right = m - 1
			}
		}
		fmt.Println(left)
	}
}
func check(m int) bool {
	for i := 0; i < len(numlist); i++ {
		count += numlist[i] / m
	}
	return count >= k
}
