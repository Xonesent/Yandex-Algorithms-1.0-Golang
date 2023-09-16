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

	k, _ := strconv.Atoi(params[0])
	text := params[1]

	var points int
	count := 0

	for i := 0; i < len(text)-k; i++ {
		for k+i < len(text) && text[i] == text[k+i] {
			count++
			i++
		}
		for count > 0 {
			points += count
			count--
		}
	}
	fmt.Println(points)
}
