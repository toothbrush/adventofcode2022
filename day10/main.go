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
	X                  int
	cycle              int
	interestingSignals []int
}

type Instruction struct {
	name string
	args []string
}

func interestingCycle(c int) bool {

	if (c-20)%40 == 0 {
		return true
	}
	return false
}

func (s *State) bumpClockReturnInterestingSignal() int {
	signal := 0
	if interestingCycle(s.cycle) {
		signal = s.cycle * s.X
		s.interestingSignals = append(s.interestingSignals, signal)
		fmt.Printf("[cycle % 3d] signal = %d\n", s.cycle, signal)
	}

	// actually bump the clock as we were asked to:
	s.cycle++

	return signal
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
		s.bumpClockReturnInterestingSignal()
	case "addx":
		arg1, err := strconv.Atoi(i.args[0])
		if err != nil {
			return err
		}
		s.bumpClockReturnInterestingSignal()
		s.bumpClockReturnInterestingSignal()
		s.X += arg1
	}
	fmt.Printf("Executing `%s`. State = %v\n", i.name, s)
	return nil
}

func NewState() State {
	s := State{}
	s.X = 1
	s.cycle = 1
	s.interestingSignals = []int{}
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
	fmt.Printf("total of %d interesting signals = %d\n",
		len(state.interestingSignals),
		sum(state.interestingSignals))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
