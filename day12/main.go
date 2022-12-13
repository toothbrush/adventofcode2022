package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

type Tile struct {
	elevation int
}

type Grid struct {
	size    int
	start_x int
	start_y int
	end_x   int
	end_y   int
	tilez   [][]Tile
}

func (g Grid) String() string {

	s := "\n"
	for i := range g.tilez {
		for j := range g.tilez[i] {
			s += fmt.Sprintf("%d", g.tilez[i][j].elevation)
		}
		s += "\n"
	}
	s += "\n"
	return s
}

func NewGrid(size int) Grid {
	g := Grid{}
	g.size = size
	g.tilez = make([][]Tile, size)
	for i := range g.tilez {
		g.tilez[i] = make([]Tile, size)
	}
	return g
}

func (g *Grid) populateGrid(lines []string) error {
	for i, l := range lines {
		if g.size != len(lines) || g.size != len(l) {
			return fmt.Errorf("eish, dimensions aren't square")
		}
		for j, c := range l {
			if c == 'S' {
				g.tilez[i][j].elevation = 'a' - 1
				g.start_x = i
				g.start_y = j
			} else if c == 'E' {
				g.tilez[i][j].elevation = 'z' + 1
				g.end_x = i
				g.end_y = j
			} else {
				// elev := int(c)
				bs := make([]byte, 4)
				if utf8.EncodeRune(bs, c) > 0 {
					g.tilez[i][j].elevation = int(bs[0])
				} else {
					return fmt.Errorf("something broken about rune.. %v", c)
				}
			}
		}
	}
	return nil
}

func run() (err error) {
	fmt.Printf("welcome to hills\n")

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

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
