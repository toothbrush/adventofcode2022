package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type Section struct {
	begin int
	end   int
}

func (s Section) fullyContains(t Section) bool {
	return s.begin <= t.begin && s.end >= t.end
}

func toAssignment(s string) (Section, error) {
	split := strings.Split(s, "-")
	if len(split) != 2 {
		return Section{}, fmt.Errorf("did not find begin-end pair: %s", s)
	}
	begin, err := strconv.Atoi(split[0])
	if err != nil {
		return Section{}, err
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		return Section{}, err
	}

	return Section{begin: begin, end: end}, nil
}

func run() error {
	s := bufio.NewScanner(os.Stdin)
	var t string
	var assignments []string

	var err error
	var x Section
	var y Section

	overlaps := 0

	for s.Scan() {
		t = s.Text()

		assignments = strings.Split(t, ",")
		if len(assignments) != 2 {
			return fmt.Errorf("did not find 2 assignments: %s", t)
		}

		x, err = toAssignment(assignments[0])
		if err != nil {
			return err
		}
		y, err = toAssignment(assignments[1])
		if err != nil {
			return err
		}

		fmt.Printf("%v, %v\n", x, y)

		if x.fullyContains(y) || y.fullyContains(x) {
			overlaps++
		}
	}

	fmt.Printf("overlaps = %d\n", overlaps)

	return nil
}
