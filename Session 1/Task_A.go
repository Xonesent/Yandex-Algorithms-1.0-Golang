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

	z := strings.Fields(params[0])
	aInt, _ := strconv.Atoi(z[0])
	bInt, _ := strconv.Atoi(z[1])

	str2 := params[1]

	min := aInt
	max := bInt
	if bInt < aInt {
		min = bInt
		max = aInt
	}
	if str2 == "freeze" {
		fmt.Print(min)
	}
	if str2 == "heat" {
		fmt.Print(max)
	}
	if str2 == "auto" {
		fmt.Print(bInt)
	}
	if str2 == "fan" {
		fmt.Print(aInt)
	}
}
