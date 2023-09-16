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
	a := make(map[string]string)
	for i := 1; i <= n; i++ {
		words := strings.Fields(params[i])
		a[words[0]] = words[1]
		a[words[1]] = words[0]
	}

	fmt.Println(a[params[n+1]])
}
