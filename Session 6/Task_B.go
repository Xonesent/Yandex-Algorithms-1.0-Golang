package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var n int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	nk := strings.Fields(params[0])
	n, _ = strconv.Atoi(string(nk[0]))
	k, _ := strconv.Atoi(string(nk[1]))
	var massiv []int
	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(strings.Fields(params[1])[i])
		massiv = append(massiv, value)
	}

	for i := 0; i < k; i++ {
		value, _ := strconv.Atoi(strings.Fields(params[2])[i])
		fmt.Println(poisk(massiv, value))
	}
}

func poisk(massiv []int, znak int) int {
	if znak < massiv[0] {
		return massiv[0]
	} else if znak > massiv[n-1] {
		return massiv[n-1]
	}

	left := 0
	right := n - 1
	for left < right {
		m := (left + right) / 2
		if znak == massiv[m] {
			return znak
		} else if znak > massiv[m] {
			left = m + 1
		} else if znak < massiv[m] {
			right = m
		}
	}
	if znak-massiv[left-1] > massiv[left]-znak {
		return massiv[left]
	} else {
		return massiv[left-1]
	}
}
