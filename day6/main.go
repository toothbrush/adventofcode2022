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
	return len(counts) == MESSAGE_LENGTH
}

// for puzzle 1 it's 4, for puzzle 2 it's 14
const MESSAGE_LENGTH = 14

func run() (err error) {
	s := bufio.NewScanner(os.Stdin)
	var t string

	for s.Scan() {
		t = s.Text()
		fmt.Printf("processing: %s...\n", t)
		ring := make([]rune, MESSAGE_LENGTH)
		for i, c := range t {
			// poorman's ring buffer
			ring[i%MESSAGE_LENGTH] = c
			// no point checking if we haven't seen enough characters yet
			if i >= MESSAGE_LENGTH {
				if all_different(ring) {
					fmt.Printf("Packet found at %d.\n", i+1)
					break
				}
			}
		}
	}

	return nil
}
