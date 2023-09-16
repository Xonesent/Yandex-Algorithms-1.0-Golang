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
	var count int
	keys := []int{}
	nk := strings.Fields(params[0])
	k, _ := strconv.Atoi(nk[1])

	list := strings.Fields(params[1])
	listint := []int{}

	for i := 0; i < len(list); i++ {
		value, _ := strconv.Atoi(list[i])
		listint = append(listint, value)
	}

	for _, value := range listint {
		_, inMap := a[value]
		if !inMap {
			a[value] = 1
		} else {
			a[value] += 1
		}
	}

	for key := range a {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	r := 0
	points := 0

	for l, value := range keys {
		for r < len(keys) && keys[r] <= value*k {
			if a[keys[r]] >= 2 {
				points++
			}
			r += 1
		}
		diff := r - 1 - l

		if a[value] >= 2 {
			count += diff * 3
			points--
		}

		count += points * 3

		if diff > 0 {
			count += (diff) * (diff - 1) * 3
		}

		if a[value] >= 3 {
			count += 1
		}
	}
	fmt.Println(count)
}
