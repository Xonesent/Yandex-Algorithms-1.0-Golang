package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var count int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ := strconv.Atoi(params[0])
	a := make(map[int]bool)
	for i := 1; i <= n; i++ {
		a[i] = false
	}
	for i := 1; i <= n; i++ {
		bc := strings.Fields(params[i])
		b, _ := strconv.Atoi(bc[0])
		c, _ := strconv.Atoi(bc[1])
		check(n, b, c, a)
	}
	fmt.Print(count)
}

func check(a int, b int, c int, d map[int]bool) {
	if b+c == a-1 && d[b+1] == false && b >= 0 && c >= 0 {
		count++
		d[b+1] = true
	}
}
