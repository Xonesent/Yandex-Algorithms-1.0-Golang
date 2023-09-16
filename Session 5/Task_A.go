package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var ans []int

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[2])

	nlist := strings.Fields(params[1])
	mlist := strings.Fields(params[3])
	ncolours := []int{}
	mcolours := []int{}

	for _, colour := range nlist {
		colour, _ := strconv.Atoi(colour)
		ncolours = append(ncolours, colour)
	}
	for _, colour := range mlist {
		colour, _ := strconv.Atoi(colour)
		mcolours = append(mcolours, colour)
	}

	ans = []int{ncolours[0], mcolours[0]}
	diff := max(ncolours[0], mcolours[0]) - min(ncolours[0], mcolours[0])

	i := 0
	idx := 0

	ans = merge(diff, i, idx, n, m, ncolours, mcolours, ans)

	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
	}
}

func merge(diff int, i int, idx int, n int, m int, ncolours []int, mcolours []int, ans []int) (anslist []int) {

	if i < n && idx < m && ncolours[i] > mcolours[idx] {
		if ncolours[i]-mcolours[idx] < diff {
			diff = ncolours[i] - mcolours[idx]
			ans = []int{ncolours[i], mcolours[idx]}
		}
		idx++
		return merge(diff, i, idx, n, m, ncolours, mcolours, ans)
	} else if i < n && idx < m && ncolours[i] < mcolours[idx] {
		if mcolours[idx]-ncolours[i] < diff {
			diff = mcolours[idx] - ncolours[i]
			ans = []int{ncolours[i], mcolours[idx]}
		}
		i++
		return merge(diff, i, idx, n, m, ncolours, mcolours, ans)
	} else if i < n && idx < m && ncolours[i] == mcolours[idx] {
		ans = []int{ncolours[i], mcolours[idx]}
	}
	return ans
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
