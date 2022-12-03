package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findCommonalities(rucksack1 string, rucksack2 string) []rune {
	common := map[rune]bool{}

	for _, c := range rucksack1 {
		if strings.Contains(rucksack2, string(c)) {
			common[c] = true
		}
	}

	list := []rune{}
	for k := range common {
		list = append(list, k)
	}

	return list
}

func priority(c rune) int {
	i := int(c)
	switch unicode.IsUpper(c) {
	case false:
		return i - 96
	case true:
		return i - 38
	}
	return -1
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	var t string
	var rucksack1 string
	var rucksack2 string
	var items int

	var totalPriority int
	totalPriority = 0

	for s.Scan() {
		t = s.Text()

		items = len(t)
		if items%2 != 0 {
			panic("non-even string length")
		}

		rucksack1 = string(t[:int(items/2)])
		rucksack2 = string(t[int(items/2):])

		fmt.Printf("%s // %s\n", rucksack1, rucksack2)
		commonalities := findCommonalities(rucksack1, rucksack2)
		if len(commonalities) != 1 {
			fmt.Println(commonalities)
			panic("Found more than 1 common element!")
		}

		fmt.Printf("%s, priority %d\n", string(commonalities[0]), priority(commonalities[0]))
		totalPriority += priority(commonalities[0])
	}
	fmt.Printf("total priority = %d\n", totalPriority)
}
