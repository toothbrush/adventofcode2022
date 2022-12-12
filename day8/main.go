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
	height       int
	is_visible   bool
	scenic_score int
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

func (g Grid) ScenicString() string {

	s := "\n"
	for i := range g.treez {
		for j := range g.treez[i] {
			s += fmt.Sprintf("% 3d ", g.treez[i][j].scenic_score)
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
	dx             int
	dy             int
	occlusion_free bool
	scenic_score   int
}
type Directions []Direction

func (directions Directions) anyVisible() bool {
	res := false
	for _, d := range directions {
		res = res || d.occlusion_free
	}
	return res
}

func NewDirections() Directions {
	directions := Directions{
		{dx: -1, dy: 0, occlusion_free: true, scenic_score: 0},
		{dx: 0, dy: -1, occlusion_free: true, scenic_score: 0},
		{dx: 0, dy: 1, occlusion_free: true, scenic_score: 0},
		{dx: 1, dy: 0, occlusion_free: true, scenic_score: 0},
	}
	return directions
}

func (d Direction) String() string {
	if d.dx == 0 && d.dy == 1 {
		return "RIGHT"
	}
	if d.dx == 0 && d.dy == -1 {
		return "LEFT"
	}
	if d.dx == 1 && d.dy == 0 {
		return "DOWN"
	}
	if d.dx == -1 && d.dy == 0 {
		return "UP"
	}

	return "boom"
}

func (g *Grid) determineVisibility() {
	for i := range g.treez {
		for j := range g.treez[i] {
			// from the pov of tree i,j walk till we're at a tree >= our height.
			// set up a bunch of directions:
			directions := NewDirections()
			// look for occlusions:
			for d_i, dir := range directions {
				// figure out if we're walking horizontally or vertically
				if dir.dy == 0 {
					for x := i + dir.dx; x >= 0 && x < g.size; x += dir.dx {
						if g.treez[x][j].height >= g.treez[i][j].height {
							// occluded from this direction!
							directions[d_i].occlusion_free = false
						}
					}
				} else {
					for y := j + dir.dy; y >= 0 && y < g.size; y += dir.dy {
						if g.treez[i][y].height >= g.treez[i][j].height {
							// occluded from this direction!
							directions[d_i].occlusion_free = false
						}
					}
				}
			}
			g.treez[i][j].is_visible = directions.anyVisible()
		}
	}
}

func (g *Grid) determineBucholicIndex() {
	for i := range g.treez {
		for j := range g.treez[i] {
			// from the pov of tree i,j walk till we're at the horizon.
			// set up a bunch of directions, start by assuming the tree is visible in that direction:
			directions := NewDirections()
			// look for occlusions:
			for d_i, dir := range directions {
				scenic := 0
				// figure out if we're walking horizontally or vertically
				if dir.dy == 0 { // walking vertically
					for x := i + dir.dx; x >= 0 && x < g.size; x += dir.dx {
						if directions[d_i].occlusion_free {
							scenic += 1
							if g.treez[x][j].height >= g.treez[i][j].height {
								// stop counting now.
								directions[d_i].occlusion_free = false
							}
						}
					}
				} else {
					for y := j + dir.dy; y >= 0 && y < g.size; y += dir.dy {
						if directions[d_i].occlusion_free {
							scenic += 1
							if g.treez[i][y].height >= g.treez[i][j].height {
								// stop counting now.
								directions[d_i].occlusion_free = false
							}
						}
					}
				}
				directions[d_i].scenic_score = scenic
				fmt.Printf("row %d col %d: scenic score %d in direction %s\n", i, j, scenic, dir)
			}
			fmt.Printf("%v\n", directions)
			g.treez[i][j].scenic_score =
				directions[0].scenic_score *
					directions[1].scenic_score *
					directions[2].scenic_score *
					directions[3].scenic_score
		}
	}
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

func (g Grid) maxScenic() int {
	max := 0
	for i := range g.treez {
		for j := range g.treez[i] {
			if g.treez[i][j].scenic_score > max {
				max = g.treez[i][j].scenic_score
			}
		}
	}
	return max
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
	g.determineVisibility()
	g.determineBucholicIndex()

	fmt.Printf("%v\n", g)
	fmt.Printf("visible: %d\n", g.countVisible())
	fmt.Printf("%s\n", g.ScenicString())
	fmt.Printf("max scenic score: %d\n", g.maxScenic())

	return nil
}
