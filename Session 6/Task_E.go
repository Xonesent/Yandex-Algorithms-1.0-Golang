package main

import (
	"fmt"
)

var a, b, c int

func main() {
	fmt.Scan(&a, &b, &c)

	var left, right int
	left = 0
	right = a + b + c

	for left < right {
		var m int = (left + right) / 2
		if check(m) {
			left = m + 1
		} else {
			right = m
		}
	}
	fmt.Println(left)
}

func check(quantity int) bool {
	return (a*2+b*3+c*4+quantity*5)*2 < 7*(a+b+c+quantity)
}
