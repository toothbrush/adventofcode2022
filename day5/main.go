package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type State struct {
	crates [][]rune
}

func (s State) String() string {
	response := ""
	for i, crates := range s.crates {
		response += fmt.Sprintf("%d: ", i+1)
		for _, item := range crates {
			response += fmt.Sprintf("%s ", string(item))
		}
		response += "\n"
	}
	return response
}

func run() (err error) {
	s := bufio.NewScanner(os.Stdin)
	var t string

	crates := regexp.MustCompile("\\[[A-Z]\\]")
	moves := regexp.MustCompile("^move")

	state := State{}
	state.crates = make([][]rune, 9) // Hm, hard-coded 9 crates

	for s.Scan() {
		t = s.Text()
		if crates.MatchString(t) {
			// we're still reading crates
			itemAt := t[1]
			state.crates[0] = append([]rune{rune(itemAt)}, state.crates[0]...)
			fmt.Printf("%s\n", t)
		}

		if moves.MatchString(t) {
			// now we're reading moves
			// fmt.Printf("%s\n", t)
		}

	}

	fmt.Printf("%v\n", state)
	return nil
}
