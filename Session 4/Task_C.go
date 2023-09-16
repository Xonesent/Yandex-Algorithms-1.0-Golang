package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Fields(str)

	var max int
	var ans string
	a := make(map[string]int)

	for i := 0; i < len(params); i++ {
		if _, inMap := a[params[i]]; !inMap {
			a[params[i]] = 0
		}
		a[params[i]]++
		if a[params[i]] > max {
			max = a[params[i]]
			ans = params[i]
		} else if a[params[i]] == max && params[i] < ans {
			ans = params[i]
		}
	}
	fmt.Print(ans)
}
