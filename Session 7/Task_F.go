package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, count1 int
	fmt.Scan(&n)
	events := make([][3]int, 0)
	alive, dead := 1, 2
	for i := 1; i < n+1; i++ {
		var day1, month1, year1, day2, month2, year2 int
		fmt.Scan(&day1, &month1, &year1, &day2, &month2, &year2)
		if count(day1, month1, year1+18) < count(day2, month2, year2) {
			events = append(events, [3]int{i, count(day1, month1, year1+18), alive})
			events = append(events, [3]int{i, min(count(day1, month1, year1+80), count(day2, month2, year2)), dead})
		} else {
			count1++
		}
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] != events[j][1] {
			return events[i][1] < events[j][1]
		}
		return events[i][2] > events[j][2]
	})
	n -= count1
	peoplealive := make(map[int]int)
	flag := false
	for i := 0; i < 2*n; i++ {
		person, action := events[i][0], events[i][2]
		if action == alive {
			peoplealive[person] = 1
			flag = true
		}
		if action == dead {
			if flag {
				for value := range peoplealive {
					fmt.Printf("%d ", value)
				}
				fmt.Printf("\n")
				flag = false
			}
			delete(peoplealive, person)
		}
	}
	if n == 0 {
		fmt.Print(0)
	}
}

func visokosniygod(year int) bool {
	var return_value bool
	if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
		return_value = false
	} else {
		return_value = true
	}
	return return_value
}

func count(day, month, year int) int {
	months := [13]int{0, 31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	count1 := 0
	for i := 1; i < year; i++ {
		if visokosniygod(i) {
			count1 += 366
		} else {
			count1 += 365
		}
	}

	for m := 1; m < month; m++ {
		if m == 2 {
			if visokosniygod(year) {
				count1 += 29
			} else {
				count1 += 28
			}
		} else {
			count1 += months[m]
		}
	}
	count1 += day
	return count1
}

func remove(l []int, item int) []int {
	for i, other := range l {
		if other == item {
			l = append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
