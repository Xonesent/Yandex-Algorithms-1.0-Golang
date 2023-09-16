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

	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	nm := strings.Fields(params[0])
	list1 := []int{}
	list2 := []int{}
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])
	for i := 1; i < 1+n; i++ {
		value, _ := strconv.Atoi(params[i])
		list1 = append(list1, value)
	}
	for i := 1 + n; i < 1+n+m; i++ {
		value, _ := strconv.Atoi(params[i])
		list2 = append(list2, value)
	}
	a := make(map[int]bool)
	b := make(map[int]bool)
	for _, num := range list1 {
		a[num] = true
	}
	for _, num := range list2 {
		b[num] = true
	}
	listans1 := []int{}
	listans2 := []int{}
	listans3 := []int{}

	for _, num := range list2 {
		if a[num] {
			listans1 = append(listans1, num)
		}
	}
	for _, num := range list1 {
		if !b[num] {
			listans2 = append(listans2, num)
		}
	}
	for _, num := range list2 {
		if !a[num] {
			listans3 = append(listans3, num)
		}
	}
	sort.Ints(listans1)
	sort.Ints(listans2)
	sort.Ints(listans3)
	fmt.Println(len(listans1))
	for i := 0; i < len(listans1); i++ {
		fmt.Printf("%v ", listans1[i])
	}
	fmt.Print("\n")
	fmt.Println(len(listans2))
	for i := 0; i < len(listans2); i++ {
		fmt.Printf("%v ", listans2[i])
	}
	fmt.Print("\n")
	fmt.Println(len(listans3))
	for i := 0; i < len(listans3); i++ {
		fmt.Printf("%v ", listans3[i])
	}
}
