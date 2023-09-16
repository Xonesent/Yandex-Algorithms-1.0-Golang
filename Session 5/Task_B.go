package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var sum int
var count int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	nk := strings.Fields(params[0])
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	numsstring := strings.Fields(params[1])

	i := 0
	idx := 1

	sum, _ = strconv.Atoi(numsstring[0])

	fmt.Println(summing(i, idx, numsstring, n, k, count, sum))
}

func summing(i int, idx int, numsstring []string, n int, k int, count int, sum int) (ans int) {
	if sum < k && idx < n {
		idx++
		value, _ := strconv.Atoi(numsstring[idx-1])
		sum += value
		return summing(i, idx, numsstring, n, k, count, sum)
	}
	if sum > k && i < n {
		i++
		value, _ := strconv.Atoi(numsstring[i-1])
		sum -= value
		return summing(i, idx, numsstring, n, k, count, sum)
	}
	if sum == k {
		count++
		if idx < n && i < n {
			i++
			idx++
			value1, _ := strconv.Atoi(numsstring[idx-1])
			value2, _ := strconv.Atoi(numsstring[i-1])
			sum = sum + value1 - value2
			return summing(i, idx, numsstring, n, k, count, sum)
		}
	}
	return count
}
