package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	parents := make(map[string][]string)
	people := make(map[string]bool)
	for i := 0; i < n-1; i++ {
		var child, predok string
		fmt.Scan(&child, &predok)
		parents[predok] = append(parents[predok], child)
		people[predok] = true
		people[child] = true
	}
	childcount := make(map[string]int)
	anslist := make([]string, 0)
	for key := range people {
		childcount[key] = counting(parents, childcount, key)
		anslist = append(anslist, key)
	}
	sort.Strings(anslist)
	for _, value := range anslist {
		fmt.Println(value, childcount[value])
	}
}

func counting(parents map[string][]string, childcount map[string]int, parent string) int {
	_, ok1 := childcount[parent]
	if ok1 {
		return childcount[parent]
	}
	_, ok2 := parents[parent]

	var listchildren []string
	if ok2 {
		listchildren = parents[parent]
	} else {
		return 0
	}
	var count int = len(listchildren)
	for i := 0; i < len(listchildren); i++ {
		count += counting(parents, childcount, listchildren[i])
	}
	return count
}
