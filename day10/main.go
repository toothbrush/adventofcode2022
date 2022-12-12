package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type State struct {
	X      int
	cycle  int
	pixels []string
}

func (s State) isPixelLit() bool {
	// sprite is at position X, and is 3 pixels wide.  So X, X+1, X+2.
	// CRT is drawing x-position cycle%40
	return s.X < (s.cycle%40) && (s.cycle%40) < s.X+3
}

type Instruction struct {
	name string
	args []string
}

func (s *State) bumpClockAndDraw() {
	if s.isPixelLit() {
		fmt.Printf("[cycle % 3d] signal = %d\n", s.cycle, 0)
		s.pixels = append(s.pixels, "#")
	} else {
		s.pixels = append(s.pixels, ".")
	}
	if s.cycle%40 == 0 {
		// add a line break after a line of pixels
		s.pixels = append(s.pixels, "\n")
	}

	// actually bump the clock as we were asked to:
	s.cycle++
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (s *State) execute(i Instruction) error {
	switch i.name {
	case "noop":
		s.bumpClockAndDraw()
	case "addx":
		arg1, err := strconv.Atoi(i.args[0])
		if err != nil {
			return err
		}
		s.bumpClockAndDraw()
		s.bumpClockAndDraw()
		s.X += arg1
	}
	fmt.Printf("Executing `%s`. State = %v\n", i.name, s)
	return nil
}

func NewState() State {
	s := State{}
	s.X = 1
	s.cycle = 1
	s.pixels = make([]string, 0)
	return s
}

func run() (err error) {
	fmt.Printf("welcome to CRT/CPU\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	state := NewState()
	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			spl := strings.Split(t, " ")
			i := Instruction{name: spl[0], args: spl[1:]}
			err = state.execute(i)
			if err != nil {
				return err
			}
		}
	}
	fmt.Printf("%s\n", strings.Join(state.pixels, ""))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
