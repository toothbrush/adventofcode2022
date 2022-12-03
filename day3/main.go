package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var t string
	var rucksack1 string
	var rucksack2 string
	var items int

	for s.Scan() {
		t = s.Text()

		items = len(t)
		if items%2 != 0 {
			panic("non-even string length")
		}

		rucksack1 = string(t[:int(items/2)])
		rucksack2 = string(t[int(items/2):])

		fmt.Printf("%s // %s\n", rucksack1, rucksack2)
	}
}
