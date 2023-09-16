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

	genome1 := params[0]
	genome2 := params[1]

	pairs := make(map[string]int)

	for i := 0; i < len(genome1)-1; i++ {
		pair := genome1[i : i+2]
		pairs[pair]++
	}

	similarity := 0
	for i := 0; i < len(genome2)-1; i++ {
		pair := genome2[i : i+2]
		if count, ok := pairs[pair]; ok {
			similarity += count
			for i := count; i > 0; i-- {
				pairs[pair]--
			}
		}
	}
	fmt.Println(similarity)
}
