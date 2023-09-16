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
	params := []int{}
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		params = append(params, value)
	}

	a := params[0]
	b := params[1]
	c := params[2]
	var result string
	var x int

	if a != 0 {
		x = (c*c - b) / a
	}

	if a == 0 && c*c == b {
		result = "MANY SOLUTIONS"
	} else if c >= 0 && a*x+b == c*c && a != 0 {
		result = strconv.Itoa((c*c - b) / a)
	} else {
		result = "NO SOLUTION"
	}
	fmt.Println(result)
}
