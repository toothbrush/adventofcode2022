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
	X     int
	cycle int
}

type Instruction struct {
	name string
	args []string
}

func (s *State) execute(i Instruction) error {
	switch i.name {
	case "noop":
		s.cycle++
	case "addx":
		arg1, err := strconv.Atoi(i.args[0])
		if err != nil {
			return err
		}
		s.X += arg1
		s.cycle++
		s.cycle++
	}
	fmt.Printf("Executing `%s`. State = %v\n", i.name, s)
	return nil
}

func NewState() State {
	s := State{}
	s.X = 1
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
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
