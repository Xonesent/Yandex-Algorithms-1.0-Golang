package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	a := make(map[string]map[string]int)
	namelist := []string{}
	for i := 0; i < len(params); i++ {
		npq := strings.Fields(params[i])
		if len(npq) != 3 {
			break
		}
		name := npq[0]
		product := npq[1]
		quantity, _ := strconv.Atoi(npq[2])

		_, inMap := a[name]
		if !inMap {
			a[name] = make(map[string]int)
			namelist = append(namelist, name)
		}

		_, inMap1 := a[name][product]
		if !inMap1 {
			a[name][product] = 0
		}
		a[name][product] += quantity
	}

	sort.Strings(namelist)

	for _, names := range namelist {
		fmt.Println(names + ":")
		productlist := []string{}
		for products := range a[names] {
			productlist = append(productlist, products)
		}
		sort.Strings(productlist)
		for _, products := range productlist {
			fmt.Println(products, a[names][products])
		}
	}
}
