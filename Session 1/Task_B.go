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

	switch {
	case a+b <= c:
		fmt.Print("NO")
	case b+c <= a:
		fmt.Print("NO")
	case c+a <= b:
		fmt.Print("NO")
	default:
		fmt.Print("YES")
	}

}
