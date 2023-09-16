package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = strings.Fields(scanner.Text())
	}

	list := []int{}
	for i := 0; i < len(params); i++ {
		paramsInt, _ := strconv.Atoi(params[i])
		list = append(list, paramsInt)
	}

	var count int
	for i := 1; i < len(list)-1; i++ {
		if list[i] > list[i-1] && list[i] > list[i+1] {
			count += 1
		}
	}
	fmt.Println(count)
}
