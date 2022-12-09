package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
			response += fmt.Sprintf("[%s] ", string(item))
		}
		response += "\n"
	}
	return response
}

func (s State) performMove(from int, to int) error {
	// get last item
	itm := s.crates[from][len(s.crates[from])-1]
	// drop last item
	s.crates[from] = s.crates[from][0 : len(s.crates[from])-1]
	// add item to target pile
	s.crates[to] = append(s.crates[to], itm)
	return nil
}

func run() (err error) {
	s := bufio.NewScanner(os.Stdin)
	var t string

	crates := regexp.MustCompile("\\[[A-Z]\\]")
	moves := regexp.MustCompile("^move ([0-9]+) from ([0-9]+) to ([0-9]+)")

	state := State{}
	state.crates = make([][]rune, 9) // Hm, hard-coded 9 crates

	for s.Scan() {
		t = s.Text()
		if crates.MatchString(t) {
			// we're still reading crates
			fmt.Printf("%s\n", t)
			for pos := 0; pos < len(state.crates); pos++ {
				if len(t) > 4*pos+1 {
					itemAt := t[4*pos+1]
					if regexp.MustCompile("[A-Z]").MatchString(string(itemAt)) {
						state.crates[pos] = append([]rune{rune(itemAt)}, state.crates[pos]...)
					}
				}
			}
		}

		if move := moves.FindStringSubmatch(t); len(move) > 0 {
			// now we're reading moves
			fmt.Printf("%v\n", move)
			iterations, err := strconv.Atoi(move[1])
			if err != nil {
				return err
			}
			from, err := strconv.Atoi(move[2])
			if err != nil {
				return err
			}
			to, err := strconv.Atoi(move[3])
			if err != nil {
				return err
			}
			for i := 0; i < iterations; i++ {
				fmt.Printf("moving from %d to %d #%d..\n", from, to, i)
				err = state.performMove(from-1, to-1) // eek fixup indexen
				if err != nil {
					return err
				}
			}
		}
	}

	fmt.Printf("%v\n", state)
	return nil
}
