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

	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[n+1])

	var ylist []int
	yleft := make([]int, n)
	yright := make([]int, n)

	for i := 1; i < n+1; i++ {
		values := strings.Fields(params[i])
		value, _ := strconv.Atoi(values[1])
		ylist = append(ylist, value)
	}

	for i := 1; i < n; i++ {
		if ylist[i]-ylist[i-1] <= 0 {
			yright[i] = yright[i-1]
		} else {
			yright[i] = (ylist[i] - ylist[i-1]) + yright[i-1]
		}
	}
	for i := n - 2; i > -1; i-- {
		if ylist[i]-ylist[i+1] <= 0 {
			yleft[i] = yleft[i+1]
		} else {
			yleft[i] = (ylist[i] - ylist[i+1]) + yleft[i+1]
		}
	}
	for i := n + 2; i < n+m+2; i++ {
		values := strings.Fields(params[i])
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		if value1 > value2 {
			fmt.Println(yleft[value2-1] - yleft[value1-1])
		}
		if value2 >= value1 {
			fmt.Println(yright[value2-1] - yright[value1-1])
		}
	}
}
