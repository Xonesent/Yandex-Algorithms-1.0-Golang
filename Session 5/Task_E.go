package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ := strconv.Atoi(strings.Fields(params[0])[0])
	k, _ := strconv.Atoi(strings.Fields(params[0])[1])

	trees := strings.Fields(params[1])
	lenans := n + 1

	var i int
	var l int
	var r int
	a := make(map[string]int)
	for idx := 0; idx < n; idx++ {
		_, inMap := a[trees[idx]]
		if !inMap {
			a[trees[idx]] = 1
		} else {
			a[trees[idx]] += 1
		}

		for a[trees[i]] > 1 {
			a[trees[i]] -= 1
			i += 1
		}
		if len(a) == k {
			if idx-i+1 < lenans {
				lenans = idx - i + 1
				l = i
				r = idx
				if lenans == k {
					break
				}
			}
		}
	}
	fmt.Println(l+1, r+1)
}
