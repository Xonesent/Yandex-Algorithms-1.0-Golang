package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var value int
var cycle int
var cycle2 int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ := strconv.Atoi(params[0])
	a := make(map[string]int)
	ans := []string{}
	for i := 1; cycle*1 < n; i += value + 1 {
		value, _ = strconv.Atoi(params[i])
		cycle++

		for idx := i + 1; cycle2 < value; {
			a[params[idx]] += 1
			if a[params[idx]] == n {
				ans = append(ans, params[idx])
			}
			idx++
			cycle2++
		}

		cycle2 = 0
	}
	fmt.Println(len(ans))
	for _, value := range ans {
		fmt.Println(value)
	}
	fmt.Println(len(a))
	for key := range a {
		fmt.Println(key)
	}
}
