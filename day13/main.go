package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run() (err error) {
	fmt.Printf("welcome to pairs n packets\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	lines := []string{}
	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			lines = append(lines, t)
			fmt.Println(t)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
