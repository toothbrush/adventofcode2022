package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type Tree struct {
	height     int
	is_visible bool
}

type Grid struct {
	size  int
	treez [][]Tree
}

func (g Grid) String() string {

	s := "\n"
	for i := range g.treez {
		for j := range g.treez[i] {
			s += fmt.Sprintf("%d", g.treez[i][j].height)
		}
		s += "\n"
	}
	s += "\n"
	for i := range g.treez {
		for j := range g.treez[i] {
			if g.treez[i][j].is_visible {
				s += "^"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	s += "\n"
	return s
}

func NewGrid(size int) Grid {
	g := Grid{}
	g.size = size
	g.treez = make([][]Tree, size)
	for i := range g.treez {
		g.treez[i] = make([]Tree, size)
	}
	return g
}

func (g *Grid) populateGrid(lines []string) error {
	if g.size != len(lines) || g.size != len(lines[0]) {
		return fmt.Errorf("eish, dimensions aren't square")
	}

	for i, l := range lines {
		for j, c := range l {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				return err
			}
			g.treez[i][j].height = height
		}
	}
	return nil
}

type Direction struct {
	dx         int
	dy         int
	is_visible bool
}
type Directions []Direction

func (directions Directions) anyVisible() bool {
	res := false
	for _, d := range directions {
		res = res || d.is_visible
	}
	return res
}

func (g Grid) countVisible() int {
	total := 0
	for i := range g.treez {
		for j := range g.treez[i] {
			if g.treez[i][j].is_visible {
				total += 1
			}
		}
	}
	return total
}

func run() (err error) {
	fmt.Printf("welcome to treez\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	lines := []string{}
	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			lines = append(lines, t)
			fmt.Println(t)
		}
	}
	g := NewGrid(len(lines))
	g.populateGrid(lines)

	fmt.Printf("%v\n", g)
	fmt.Printf("visible: %d\n", g.countVisible())

	return nil
}
