package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	poiskpapi := make(map[string]string)
	parents := make(map[string]bool)
	people := make(map[string]bool)
	for i := 0; i < n-1; i++ {
		var child, predok string
		fmt.Scan(&child, &predok)
		poiskpapi[child] = predok
		parents[child] = true
		people[child] = true
		people[predok] = true
	}
	anslist := make([]string, 0)
	for key := range people {
		anslist = append(anslist, key)
	}
	sort.Strings(anslist)

	glubini := make(map[string]int)
	for i := 0; i < len(anslist); i++ {
		glubini[anslist[i]] = glubina(anslist[i], parents, poiskpapi, glubini)
	}

	for i := 0; i < len(anslist); i++ {
		fmt.Println(anslist[i], glubini[anslist[i]])
	}
}

func glubina(person string, parents map[string]bool, poiskpapi map[string]string, glubini map[string]int) int {
	_, ok1 := glubini[person]
	if ok1 {
		return glubini[person]
	}
	_, ok2 := parents[person]
	if !ok2 {
		return 0
	} else {
		return glubina(poiskpapi[person], parents, poiskpapi, glubini) + 1
	}
}
