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

func run() (err error) {
	fmt.Printf("welcome to CRT/CPU\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			fmt.Println(t)
		}
	}
	return nil
}
