package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var t string
	for s.Scan() {
		t = s.Text()
		fmt.Println(t)
	}
}
