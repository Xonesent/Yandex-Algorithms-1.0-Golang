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

	a := make(map[string]int)
	n, _ := strconv.Atoi(params[0])
	ci := strings.Fields(params[1])
	k, _ := strconv.Atoi(params[2])
	pj := strings.Fields(params[3])

	for i := 0; i < k; i++ {
		a[pj[i]]--
	}

	ciIntList := []int{}
	for i := 1; i <= n; i++ {
		ciInt, _ := strconv.Atoi(ci[i-1])
		ciIntList = append(ciIntList, ciInt)
		iString := strconv.Itoa(i)
		a[iString] = a[iString] + ciIntList[i-1]
		if a[iString] < 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
