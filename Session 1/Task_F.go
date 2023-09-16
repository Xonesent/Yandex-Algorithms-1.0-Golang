package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])

	arr := []int{a, b, c, d}
	sort.Ints(arr)
	switch {
	case a == c:
		fmt.Println(a, b+d)
	case a == d:
		fmt.Println(a, b+c)
	case b == c:
		fmt.Println(b, a+d)
	case b == d:
		fmt.Println(b, a+c)
	case (arr[len(arr)-1] == a && arr[len(arr)-2] == b) || (arr[len(arr)-1] == b && arr[len(arr)-2] == a) || (arr[len(arr)-1] == c && arr[len(arr)-2] == d) || (arr[len(arr)-1] == d && arr[len(arr)-2] == c):
		fmt.Println(arr[len(arr)-2], arr[len(arr)-1]+arr[len(arr)-4])
	case (arr[len(arr)-1] == a && arr[len(arr)-3] == b) || (arr[len(arr)-1] == b && arr[len(arr)-3] == a) || (arr[len(arr)-1] == c && arr[len(arr)-3] == d) || (arr[len(arr)-1] == d && arr[len(arr)-3] == c):
		fmt.Println(arr[len(arr)-1], arr[len(arr)-3]+arr[len(arr)-4])
	case (arr[len(arr)-1] == a && arr[len(arr)-4] == b) || (arr[len(arr)-1] == b && arr[len(arr)-4] == a) || (arr[len(arr)-1] == c && arr[len(arr)-4] == d) || (arr[len(arr)-1] == d && arr[len(arr)-4] == c):
		fmt.Println(arr[len(arr)-1], arr[len(arr)-3]+arr[len(arr)-4])
	}

}
