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

	whn := strings.Fields(params[0])
	w, _ := strconv.Atoi(string(whn[0]))
	h, _ := strconv.Atoi(string(whn[1]))
	n, _ := strconv.Atoi(string(whn[2]))

	left := 0
	right := max(w, h) * n
	for left < right {
		m := (left + right) / 2
		if (m/w)*(m/h) >= n {
			right = m
		} else {
			left = m + 1
		}
	}
	fmt.Println(left)
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
