package main

import (
	"bufio"
	"fmt"
	"math"
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
		params = append(params, scanner.Text())
	}

	quantity, _ := strconv.Atoi(params[0])
	list := []int{}
	bebra := strings.Fields(params[1])
	for _, value := range bebra {
		valueInt, _ := strconv.Atoi(value)
		list = append(list, valueInt)
	}
	number, _ := strconv.Atoi(params[2])
	ans := list[0]

	for i := 1; i < quantity; i++ {
		if math.Abs(float64(list[i])-float64(number)) < math.Abs(float64(ans)-float64(number)) {
			ans = list[i]
		}
	}

	fmt.Println(ans)

}
