package main

import "fmt"

func checkValue(i, j int, mp [][]int, n, m int) int {
	if i < 0 || j < 0 || i >= n || j >= m {
		return 0
	}
	if mp[i][j] == -1 {
		return 1
	}
	return 0
}

func checkBlock(i, j int, mp [][]int, n, m int) int {
	result := 0

	result += checkValue(i-1, j-1, mp, n, m)
	result += checkValue(i-1, j, mp, n, m)
	result += checkValue(i-1, j+1, mp, n, m)

	result += checkValue(i, j-1, mp, n, m)
	result += checkValue(i, j+1, mp, n, m)

	result += checkValue(i+1, j-1, mp, n, m)
	result += checkValue(i+1, j, mp, n, m)
	result += checkValue(i+1, j+1, mp, n, m)

	return result
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	mp := make([][]int, 0)
	for i := 0; i < n; i++ {
		mp = append(mp, make([]int, m))
	}
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		mp[x-1][y-1] = -1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mp[i][j] != -1 {
				mp[i][j] = checkBlock(i, j, mp, n, m)
			}
		}
	}
	for _, line := range mp {
		for _, value := range line {
			if value == -1 {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d ", value)
			}
		}
		fmt.Println()
	}
}
