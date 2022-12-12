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
	min_dimension int

	tail_history map[string]bool
}

func NewPos() Pos {
	// remember, the tail visited the start!
	return Pos{tail_history: map[string]bool{"0,0": true}}
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

func min(nums ...int) (min int) {
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func (p Pos) headTailTouching() bool {
	// figure out where the tail may be, to be touching the head
	allowed := make(map[string]bool)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// eish, poorman's Set
			allowed[fmt.Sprintf("%d,%d", p.head_x+i, p.head_y+j)] = true
		}
	}
	fmt.Printf("%v\n", allowed)

	// if it's there, yay.
	return allowed[fmt.Sprintf("%d,%d", p.tail_x, p.tail_y)]
}

// retain the sign but make sure abs(sign(num)) == 1, unless zero.
func sign(num int) int {
	if num < 0 {
		return -1
	} else if num > 0 {
		return 1
	} else {
		return 0
	}
}

func (p *Pos) bringTailForTheRide() {
	diffx := p.head_x - p.tail_x
	diffy := p.head_y - p.tail_y

	// oh scheisse might be a diagonal move. even so, sign() helps us!
	// we can get this done by using diffx and diffy, but at most one step in those directions
	p.tail_x += sign(diffx)
	p.tail_y += sign(diffy)
}

func (p *Pos) performMove(dir string) error {
	dx, dy, err := offset(dir)
	if err != nil {
		return err
	}
	p.head_x += dx
	p.head_y += dy

	if !p.headTailTouching() {
		// scheisse gotta move
		p.bringTailForTheRide()
	}

	p.max_dimension = max(p.head_x, p.head_y, p.tail_x, p.tail_y, p.max_dimension)
	p.min_dimension = min(p.head_x, p.head_y, p.tail_x, p.tail_y, p.max_dimension)
	p.tail_history[fmt.Sprintf("%d,%d", p.tail_x, p.tail_y)] = true
	return nil
}

func PrintBoard(p Pos) {
	max_coord := p.max_dimension
	min_coord := p.min_dimension
	s := ""
	var here string
	for x := max_coord; x >= min_coord; x-- {
		for y := min_coord; y < max_coord; y++ {
			here = "."
			if p.tail_history[fmt.Sprintf("%d,%d", x, y)] {
				here = "#"
			}
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

func countVisited(history map[string]bool) (vis int) {
	for range history {
		vis++
	}
	return vis
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
	fmt.Printf("tail visited sites: %d\n", countVisited(state.tail_history))
	return nil
}
