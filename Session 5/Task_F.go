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

	a := make(map[int]int)
	n, _ := strconv.Atoi(params[0])
	k, _ := strconv.Atoi(params[2])
	var count int

	for i := 3; i < 3+k; i++ {
		key, _ := strconv.Atoi(strings.Fields(params[i])[0])
		value, _ := strconv.Atoi(strings.Fields(params[i])[1])
		_, inMap := a[value]
		if !inMap {
			a[value] = key
		} else {
			a[value] = max(key, a[value])
		}
	}

	keys := make([]int, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	condlist := make([]int, 0, n)
	for _, value := range strings.Fields(params[1]) {
		valueint, _ := strconv.Atoi(value)
		condlist = append(condlist, valueint)
	}
	sort.Ints(condlist)

	r := 0
	for _, value := range condlist {
		for r < len(a) && value > a[keys[r]] {
			r += 1
		}
		count += keys[r]
	}
	fmt.Println(count)
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
