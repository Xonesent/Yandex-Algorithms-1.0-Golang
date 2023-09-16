package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	massiv := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&massiv[i])
	}
	for i := 0; i < k; i++ {
		var value int
		fmt.Scan(&value)
		if poisk(value, massiv, n) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func poisk(znak int, massiv []int, n int) bool {
	left := -1
	right := n
	for right-left > 1 {
		m := (left + right) / 2
		if massiv[m] == znak {
			return true
		} else if massiv[m] > znak {
			right = m
		} else {
			left = m
		}
	}
	return false
}
