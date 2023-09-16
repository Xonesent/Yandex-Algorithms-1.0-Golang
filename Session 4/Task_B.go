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

	a := make(map[string]int)
	for i := 0; i < len(params); i++ {
		words := strings.Fields(params[i])
		for i := 0; i < len(words); i++ {
			if _, inMap := a[words[i]]; !inMap {
				fmt.Printf("%d ", 0)
			}
			if _, inMap := a[words[i]]; inMap {
				fmt.Printf("%d ", a[words[i]])
			}
			a[words[i]] += 1
		}

	}
}
