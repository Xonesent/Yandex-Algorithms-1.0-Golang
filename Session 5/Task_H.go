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

	a := make(map[byte]int)
	k, _ := strconv.Atoi(strings.Fields(params[0])[1])
	text := string(params[1])

	left := 0
	lenans := 0
	leftans := 0

	for idx := 0; idx < len(text); idx++ {
		_, inMap := a[text[idx]]
		if !inMap {
			a[text[idx]] = 1
		} else {
			a[text[idx]] += 1
		}
		if a[text[idx]] > k {
			for text[left] != text[idx] {
				a[text[left]] -= 1
				left++
			}
			a[text[idx]] -= 1
			left++
		}
		if idx+1-left > lenans {
			lenans = idx + 1 - left
			leftans = left + 1
		}
	}
	fmt.Println(lenans, leftans)
}
