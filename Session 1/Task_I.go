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
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])
	e, _ := strconv.Atoi(params[4])

	switch {
	case (a <= d && b <= e) || (a <= e && b <= d):
		fmt.Println("YES")
	case (a <= d && c <= e) || (a <= e && c <= d):
		fmt.Println("YES")
	case (c <= d && b <= e) || (c <= e && b <= d):
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
