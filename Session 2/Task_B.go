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
	params := []int{}
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		params = append(params, value)
	}
	params = append(params[:len(params)-1], params[len(params):]...)
	anslist := ""
	ans := ""

	for i := 0; i < (len(params) - 1); i++ {
		switch {
		case params[i] == params[i+1]:
			anslist += "CONSTANT "
		case params[i] > params[i+1]:
			anslist += "WEAKLY DESCENDING "
		case params[i] < params[i+1]:
			anslist += "WEAKLY ASCENDING "
		}
	}

	switch {
	case strings.Count(anslist, "WEAKLY ASCENDING") == len(params)-1:
		ans = "ASCENDING"
	case strings.Count(anslist, "WEAKLY DESCENDING") == len(params)-1:
		ans = "DESCENDING"
	case strings.Count(anslist, "CONSTANT") == len(params)-1:
		ans = "CONSTANT"
	case strings.Contains(anslist, "CONSTANT") && strings.Contains(anslist, "WEAKLY ASCENDING") && !strings.Contains(anslist, "WEAKLY DESCENDING"):
		ans = "WEAKLY ASCENDING"
	case strings.Contains(anslist, "CONSTANT") && strings.Contains(anslist, "WEAKLY DESCENDING") && !strings.Contains(anslist, "WEAKLY ASCENDING"):
		ans = "WEAKLY DESCENDING"
	default:
		ans = "RANDOM"
	}

	fmt.Println(ans)
}
