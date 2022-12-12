package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const KNOT_COUNT = 10 // including head

type Pos struct {
	knot_x []int
	knot_y []int

	max_dimension int
	min_dimension int

	tail_history map[string]bool
}

func NewPos() Pos {
	// remember, the tail visited the start!
	p := Pos{}
	p.tail_history = map[string]bool{"0,0": true}
	p.knot_x = make([]int, KNOT_COUNT)
	p.knot_y = make([]int, KNOT_COUNT)
	return p
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

func (p Pos) knotsTouching(head int) bool {
	// figure out where the tail may be, to be touching the head
	allowed := make(map[string]bool)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// eish, poorman's Set
			allowed[fmt.Sprintf("%d,%d", p.knot_x[head]+i, p.knot_y[head]+j)] = true
		}
	}

	// if it's there, yay.
	return allowed[fmt.Sprintf("%d,%d", p.knot_x[head+1], p.knot_y[head+1])]
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

func (p *Pos) bringTailForTheRide(head int) {
	diffx := p.knot_x[head] - p.knot_x[head+1]
	diffy := p.knot_y[head] - p.knot_y[head+1]

	// oh scheisse might be a diagonal move. even so, sign() helps us!
	// we can get this done by using diffx and diffy, but at most one step in those directions
	p.knot_x[head+1] += sign(diffx)
	p.knot_y[head+1] += sign(diffy)
}

func (p *Pos) performMove(dir string) error {
	dx, dy, err := offset(dir)
	if err != nil {
		return err
	}
	p.knot_x[0] += dx
	p.knot_y[0] += dy

	for i := 0; i < KNOT_COUNT-1; i++ {
		if !p.knotsTouching(i) {
			// scheisse gotta move
			p.bringTailForTheRide(i)
		}
	}

	p.max_dimension = max(p.knot_x[0], p.knot_y[0], p.max_dimension)
	p.min_dimension = min(p.knot_x[0], p.knot_y[0], p.max_dimension)
	p.tail_history[fmt.Sprintf("%d,%d", p.knot_x[KNOT_COUNT-1], p.knot_y[KNOT_COUNT-1])] = true
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
			for i := KNOT_COUNT - 1; i > 0; i-- {
				if x == p.knot_x[i] && y == p.knot_y[i] {
					here = fmt.Sprintf("%d", i)
				}
			}
			if x == p.knot_x[0] && y == p.knot_y[0] {
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
