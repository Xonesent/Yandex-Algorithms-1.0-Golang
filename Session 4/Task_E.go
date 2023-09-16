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

	n, _ := strconv.Atoi(params[0])
	ans := 0
	a := make(map[int]int)
	for i := 1; i <= n; i++ {
		wh := strings.Fields(params[i])
		w, _ := strconv.Atoi(wh[0])
		h, _ := strconv.Atoi(wh[1])

		_, inMap := a[w]
		if !inMap {
			a[w] = h
		} else if a[w] < h {
			a[w] = h
		}
	}
	for _, value := range a {
		ans += value
	}
	fmt.Println(ans)
}
