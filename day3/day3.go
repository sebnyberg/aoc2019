package day3

import (
	"fmt"
	"strconv"

	"github.com/sebnyberg/aoc2019/util"
)

type WireGrid struct {
	FirstWire  []GridLine
	SecondWire []GridLine
	Port       Point
}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

type Point struct {
	X int
	Y int
}

type GridLine struct {
	Start     Point
	End       Point
	Direction Direction
}

func (l GridLine) CrossesLine(l2 GridLine) *Point {
	// Two lines going in the same direction can't cross each other (in this task)
	if l.Direction == l2.Direction {
		return nil
	}

	// If the first line is horizontal, check its y-coordinate against
	// the y-coordinate interval of the other line
	if l.Direction == Horizontal {
		if l.Start.Y < l2.Start.Y && l.Start.Y > l2.End.Y ||
			l.Start.Y > l2.Start.Y && l.Start.Y < l2.End.Y {
			return &Point{
				X: l2.Start.X,
				Y: l.Start.Y,
			}
		}
	} else if l.Direction == Vertical {
		if l.Start.X < l2.Start.X && l.Start.X > l2.End.X ||
			l.Start.X > l2.Start.X && l.Start.X < l2.End.X {
			return &Point{
				X: l.Start.X,
				Y: l2.Start.Y,
			}
		}
	}

	return nil
}

func NewGrid() WireGrid {
	return WireGrid{
		FirstWire:  nil,
		SecondWire: nil,
		Port: Point{
			X: 0,
			Y: 0,
		},
	}
}

func (g *WireGrid) PutWire(wiring []string) {
	start := g.GetPort()
	g.Grid[start.X][start.Y]++

	for _, part := range wiring {
		direction := string(part[0])
		steps, err := strconv.Atoi(string(part[1:]))
		util.CheckErr(err)
		if steps <= 1 {
			panic("few steps LOL")
		}

		switch direction {
		case "U":
			for i := 1; i <= steps; i++ {
				g.Grid[start.X][start.Y+i]++
			}
			start = Point{start.X, start.Y + steps}
		case "D":
			for i := 1; i <= steps; i++ {
				g.Grid[start.X][start.Y-i]++
			}
			start = Point{start.X, start.Y - steps}
		case "R":
			for i := 1; i <= steps; i++ {
				g.Grid[start.X+i][start.Y]++
			}
			start = Point{start.X + steps, start.Y}
		case "L":
			for i := 1; i <= steps; i++ {
				g.Grid[start.X-i][start.Y]++
			}
			start = Point{start.X - steps, start.Y}
		default:
			panic(fmt.Sprintf("unrecognized direction: %v", direction))
		}
	}
}

func (g *WireGrid) GetPort() Point {
	return Point{GRID_SIZE / 2, GRID_SIZE / 2}
}

func (g *WireGrid) GetCrossings() []Point {
	crossings := []Point{}

	for i := range g.Grid {
		for j := range g.Grid[i] {
			if g.Grid[i][j] > 1 {
				if i == j && i == GRID_SIZE/2 {
					continue
				}
				crossings = append(crossings, Point{
					X: i,
					Y: j,
				})
			}
		}
	}
	return crossings
}

func (p Point) DistanceTo(p2 Point) int {
	xDist := p.X - p2.X
	if xDist < 0 {
		xDist = -xDist
	}

	yDist := p.Y - p2.Y
	if yDist < 0 {
		yDist = -yDist
	}

	return xDist + yDist
}
