package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Dict = map[string]int

func main() {
	result := 0
	file, _ := ioutil.ReadFile("input.txt")
	str := string(file)
	records := strings.Split(str, "\n")

	params := strings.Fields(records[0])
	w := records[1]
	s := records[2]
	wordLen, _ := strconv.Atoi(params[0])
	strLen, _ := strconv.Atoi(params[1])

	wDict := makeDict(w, wordLen)
	sDict := makeDict(s, wordLen)
	matchLetters := matchDict(wDict, sDict)
	if matchLetters == len(wDict) {
		result++
	}

	for i := wordLen; i < strLen; i++ {
		matchLetters += modifyDict(sDict, wDict, string(s[i-wordLen]), -1)
		matchLetters += modifyDict(sDict, wDict, string(s[i]), +1)
		if matchLetters == len(wDict) {
			result++
		}
	}

	fmt.Println(result)
}

func makeDict(str string, lenW int) Dict {
	word := make(Dict)

	for i := 0; i < lenW; i++ {
		_, ok := word[string(str[i])]
		if !ok {
			word[string(str[i])] = 0
		}

		word[string(str[i])] += 1
	}
	return word
}

func matchDict(map1 Dict, map2 Dict) int {
	matches := 0
	for sym := range map1 {
		if map2[sym] > 0 && map1[sym] == map2[sym] {
			matches++
		}
	}
	return matches
}

func modifyDict(sDict Dict, wDict Dict, symbol string, countModifier int) int {
	ans := 0

	if wDict[symbol] > 0 && sDict[symbol] == wDict[symbol] {
		ans = -1
	}

	sDict[symbol] += countModifier

	if wDict[symbol] > 0 && wDict[symbol] == sDict[symbol] {
		ans = 1
	}

	return ans

}
