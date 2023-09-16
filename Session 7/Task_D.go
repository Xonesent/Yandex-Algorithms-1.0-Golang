package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, count int
	fmt.Scan(&n)
	events := make([][3]int, 0)
	vxod, vixod := 1, 2
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		if end-start >= 5 {
			events = append(events, [3]int{i, start, vxod})
			events = append(events, [3]int{i, end - 5, vixod})
		} else {
			count++
		}
	}
	n -= count
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][2] < events[j][2]
	})
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	bestsum, best1, best2 := 0, 0, 5
	customers := make(map[int]int)

	for i := 0; i < 2*n; i++ {
		customer, time, action := events[i][0], events[i][1], events[i][2]
		if action == vxod {
			customers[customer] = time
			if len(customers) > bestsum {
				best1 = time
				best2 = time + 5
				bestsum = len(customers)
			}
		}
		secondcount := 0
		for idx := i + 1; idx < 2*n; idx++ {
			customer2, time2, action2 := events[idx][0], events[idx][1], events[idx][2]
			_, ok := customers[customer2]
			if action2 == vxod {
				secondcount += 1
			}
			if time2-time >= 5 && len(customers)+secondcount > bestsum {
				best1 = time
				best2 = time2
				bestsum = len(customers) + secondcount
			}
			if action2 == vixod && !ok {
				secondcount -= 1
			}
		}
		if action == vixod {
			delete(customers, customer)
		}
	}

	fmt.Printf("%d %d %d", bestsum, best1, best2)
}
