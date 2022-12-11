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

	return nil
}
