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
	a := make(map[int]bool)
	for i := 1; i < n+1; i++ {
		xy := strings.Fields(params[i])
		x, _ := strconv.Atoi(xy[0])
		a[x] = true
	}
	fmt.Println(len(a))
}
