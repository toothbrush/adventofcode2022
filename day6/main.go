package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func all_different(ring []rune) bool {
	counts := make(map[rune]int)
	for _, c := range ring {
		counts[c]++
	}
	// fmt.Printf("%v\n", counts)
	return len(counts) == 4
}

func run() (err error) {
	s := bufio.NewScanner(os.Stdin)
	var t string

	for s.Scan() {
		t = s.Text()
		fmt.Printf("processing: %s...\n", t)
		ring := make([]rune, 4)
		for i, c := range t {
			// poorman's ring buffer
			ring[i%4] = c
			// no point checking if we haven't seen enough characters yet
			if i >= 4 {
				if all_different(ring) {
					fmt.Printf("Packet found at %d.\n", i+1)
					break
				}
			}
		}
	}

	return nil
}
