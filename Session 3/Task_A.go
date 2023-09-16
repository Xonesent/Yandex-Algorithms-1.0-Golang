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
	a := make(map[string]int)
	for _, key := range params {
		a[key] = 1
	}
	fmt.Println(len(a))
}
