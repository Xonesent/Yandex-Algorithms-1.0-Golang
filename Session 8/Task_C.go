package main

import (
	"fmt"
	"sort"
)

func main() {
	var value int = 1
	values := make([]int, 0)
	valuesmap := make(map[int]bool)
	for {
		fmt.Scan(&value)
		if value == 0 {
			break
		}
		if !valuesmap[value] {
			valuesmap[value] = true
			values = append(values, value)
		}
	}
	sort.Ints(values)
	fmt.Println(values[len(values)-2])
}

// :) xdd
