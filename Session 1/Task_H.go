package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	n, _ := strconv.Atoi(params[2])
	m, _ := strconv.Atoi(params[3])

	if Max(MinTime(a, n), MinTime(b, m)) <= Min(MaxTime(a, n), MaxTime(b, m)) {
		fmt.Println(Max(MinTime(a, n), MinTime(b, m)), Min(MaxTime(a, n), MaxTime(b, m)))
	} else {
		fmt.Println(-1)
	}
}

func MinTime(a int, b int) int {
	return b + (b-1)*a
}

func MaxTime(a int, b int) int {
	return b + (b+1)*a
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
