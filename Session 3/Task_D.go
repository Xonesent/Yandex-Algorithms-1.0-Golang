package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	a := make(map[string]bool)
	for i := 0; i < len(params); i++ {
		words := strings.Fields(params[i])
		for i := 0; i < len(words); i++ {
			a[words[i]] = true
		}
	}
	fmt.Println(len(a))
}
