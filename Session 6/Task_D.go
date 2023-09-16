package main

import (
	"fmt"
)

var n, a, b, w, h int

func main() {
	fmt.Scan(&n, &a, &b, &w, &h)
	l := -1
	r := max(w, h) + 1

	for r-l > 1 {
		m := (l + r) / 2
		d2 := m * 2
		side_a := a + d2
		side_b := b + d2
		s_ab := side_a * side_b

		success := check(side_a, side_b, s_ab)
		if !success {
			success = check(side_b, side_a, s_ab)
		}

		if success {
			l = m
		} else {
			r = m
		}
	}
	if l > -1 {
		fmt.Println(l)
	} else {
		fmt.Println(0)
	}
}

func check(side1 int, side2 int, s_ab int) bool {
	x := w / side1
	y := h / side2
	count_ab := x * y
	return count_ab >= n && s_ab*count_ab <= w*h
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
