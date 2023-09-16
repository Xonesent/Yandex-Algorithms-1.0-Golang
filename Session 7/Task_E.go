package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, count1, kassi, count, ans int
	var flag bool = false
	fmt.Scan(&n)
	kassi = n
	events := make([][3]int, 0)
	vxod, vixod := 1, 2
	for i := 0; i < n; i++ {
		var starth, startm, endh, endm int
		fmt.Scan(&starth, &startm, &endh, &endm)
		if starth*60+startm == endh*60+endm {
			events = append(events, [3]int{i, 0, vxod})
			events = append(events, [3]int{i, 24 * 60, vixod})
		} else if starth*60+startm > endh*60+endm {
			events = append(events, [3]int{i, 0, vxod})
			events = append(events, [3]int{i, endh*60 + endm, vixod})
			events = append(events, [3]int{i, starth*60 + startm, vxod})
			events = append(events, [3]int{i, 24 * 60, vixod})
			count1 += 1
		} else {
			events = append(events, [3]int{i, starth*60 + startm, vxod})
			events = append(events, [3]int{i, endh*60 + endm, vixod})
		}
	}
	n += count1
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][2] < events[j][2]
	})
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	var starttime int
	for i := 0; i < 2*n; i++ {
		time, action := events[i][1], events[i][2]
		if action == vxod {
			count += 1
		}
		if (!flag) && (count == kassi) {
			starttime = time
			flag = true
		}
		if action == vixod {
			count -= 1
		}
		if (flag) && (count != kassi) {
			ans += time - starttime
			flag = false
		}
	}
	fmt.Printf("%d", ans)
}
