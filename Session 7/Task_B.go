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

	nm := strings.Fields(params[0])
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	events := make([][2]int, 0)
	left, right, point := 1, 3, 2
	for i := 1; i <= n; i++ {
		se := strings.Fields(params[i])
		start, _ := strconv.Atoi(se[0])
		end, _ := strconv.Atoi(se[1])
		if start > end {
			start, end = end, start
		}
		events = append(events, [2]int{start, left})
		events = append(events, [2]int{end, right})
	}
	mpoints := strings.Fields(params[n+1])
	for i := 0; i < m; i++ {
		tick, _ := strconv.Atoi(mpoints[i])
		events = append(events, [2]int{tick, point})
	}
	sort.SliceStable(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	sort.SliceStable(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	var count int
	ans := make(map[int]int)
	for i := 0; i < len(events); i++ {
		if events[i][1] == left {
			count++
		} else if events[i][1] == point {
			ans[events[i][0]] = count
		} else if events[i][1] == right {
			count--
		}
	}
	for i := 0; i < m; i++ {
		tick, _ := strconv.Atoi(mpoints[i])
		fmt.Print(ans[tick], " ")
	}
}
