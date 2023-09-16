package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	events := make([][2]int, 0)
	left, right := 1, 2
	for i := 0; i < m; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		events = append(events, [2]int{start, left})
		events = append(events, [2]int{end, right})
	}
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	var count, prevcount, ans int
	for i := 0; i < len(events); i++ {
		if count == 0 {
			ans += events[i][0] - prevcount
		}
		if events[i][1] == left {
			count++
		} else if events[i][1] == right {
			count--
		}
		if count == 0 {
			prevcount = events[i][0] + 1
		}
	}
	ans += n - prevcount
	fmt.Println(ans)
}
