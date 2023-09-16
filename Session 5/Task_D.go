package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var count int
var idx int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	nr := strings.Fields(params[0])
	n, _ := strconv.Atoi(nr[0])
	r, _ := strconv.Atoi(nr[1])
	prefstr := strings.Fields(params[1])
	pref := []int{}
	for i := 0; i < len(prefstr); i++ {
		value, _ := strconv.Atoi(prefstr[i])
		pref = append(pref, value)
	}

	idx = 0

	for i := 0; i < n; i++ {
		for idx < n && pref[idx]-pref[i] <= r {
			idx++
		}
		count += n - idx
	}

	fmt.Println(count)
}
