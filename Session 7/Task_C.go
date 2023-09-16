package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	ansmap := make(map[int]int)
	students := make([]int, n)
	events := make([][]int, 0)

	for i := 0; i < n; i++ {
		fmt.Scan(&students[i])
	}

	for i := 0; i < n; i++ {
		events = append(events, []int{students[i], 1})
		events = append(events, []int{students[i] + d, 2})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	var max, count int
	for i := 0; i < 2*n; i++ {
		if events[i][1] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count--
		}
	}
	fmt.Println(max)
	variant := 1

	for i := 0; i < 2*n; i++ {
		if events[i][1] == 1 {
			ansmap[events[i][0]] = variant
			if variant < max {
				variant++
			} else {
				variant = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(ansmap[students[i]], " ")
	}
}
