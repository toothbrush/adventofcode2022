package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run() (err error) {
	fmt.Printf("welcome to monkeys\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			fmt.Printf("%s\n", t)
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
