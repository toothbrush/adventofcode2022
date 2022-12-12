package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	head_x int
	head_y int
	tail_x int
	tail_y int

	max_dimension int

	tail_history map[string]bool
}

func NewPos() Pos {
	return Pos{tail_history: make(map[string]bool)}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func offset(dir string) (dx int, dy int, err error) {
	switch dir {
	case "U":
		return 1, 0, nil
	case "D":
		return -1, 0, nil
	case "L":
		return 0, -1, nil
	case "R":
		return 0, 1, nil
	}

	return 0, 0, fmt.Errorf("unknown direction: %s", dir)
}

func max(nums ...int) int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func (p *Pos) performMove(dir string) error {
	dx, dy, err := offset(dir)
	if err != nil {
		return err
	}
	p.head_x += dx
	p.head_y += dy

	if p.max_dimension < max(p.head_x, p.head_y, p.tail_x, p.tail_y) {
		p.max_dimension = max(p.head_x, p.head_y, p.tail_x, p.tail_y)
	}
	p.tail_history[fmt.Sprintf("%d,%d", p.tail_x, p.tail_y)] = true
	return nil
}

func PrintBoard(p Pos) {
	max_coord := p.max_dimension
	s := ""
	var here string
	for x := max_coord; x >= 0; x-- {
		for y := 0; y < max_coord; y++ {
			here = "."
			if x == 0 && y == 0 {
				here = "s"
			}
			if x == p.tail_x && y == p.tail_y {
				here = "T"
			}
			if x == p.head_x && y == p.head_y {
				here = "H"
			}
			s += here

		}
		s += "\n"
	}
	fmt.Println(s)
}

func run() (err error) {
	fmt.Printf("welcome to rope\n")

	state := NewPos()
	s := bufio.NewScanner(os.Stdin)
	var t string

	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			fmt.Println(t)
			spl := strings.Split(t, " ")
			rpt, err := strconv.Atoi(spl[1])
			if err != nil {
				return err
			}
			for i := 0; i < rpt; i++ {
				err = state.performMove(spl[0])
				if err != nil {
					return err
				}
			}
		}
	}
	PrintBoard(state)
	fmt.Printf("final state: %v\n", state)
	return nil
}
