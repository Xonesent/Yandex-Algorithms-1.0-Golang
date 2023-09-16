package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ := strconv.Atoi(params[0])
	var count int

	for i := 1; i <= n; i++ {
		a := []int{}
		b := make(map[[2]int]bool)
		x1y1 := strings.Fields(string(params[i]))
		x1, _ := strconv.Atoi(x1y1[0])
		y1, _ := strconv.Atoi(x1y1[1])
		for idx := 1; idx <= n; idx++ {
			x2y2 := strings.Fields(string(params[idx]))
			x2, _ := strconv.Atoi(x2y2[0])
			y2, _ := strconv.Atoi(x2y2[1])
			x := x1 - x2
			y := y1 - y2
			_, ok := b[[2]int{-x, -y}]
			if ok {
				count--
			}
			b[[2]int{x, y}] = true
			a = append(a, x*x+y*y)
		}
		sort.Ints(a)
		var right int
		for index := 0; index < len(a)-1; index++ {
			for right < len(a) && a[right] == a[index] {
				right++
			}
			count += right - index - 1
		}
	}
	fmt.Println(count)
}
