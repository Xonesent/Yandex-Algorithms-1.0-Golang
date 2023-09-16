package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	xyz := strings.Fields(params[0])
	x := xyz[0]
	y := xyz[1]
	z := xyz[2]
	var count int
	number := params[1]
	b := []string{x, y, z}
	c := make(map[string]int)
	for _, value := range number {
		c[string(value)] = 1
	}
	for key := range c {
		for i := 0; i < len(b); i++ {
			if key == b[i] {
				count++
			}
		}
	}
	fmt.Println(len(c) - count)
}
