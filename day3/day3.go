package day3

import (
	"fmt"
	"strconv"

	"github.com/sebnyberg/aoc2019/util"
)

type Wire []GridLine

type LineType int

const (
	Horizontal LineType = iota
	Vertical
)

type Point struct {
	X int
	Y int
}
type Intersection struct {
	Point
	wireLength int
}

type GridLine struct {
	Start    Point
	End      Point
	LineType LineType
}

func (l GridLine) Distance() int {
	return l.Start.DistanceTo(l.End)
}

func WireLength(ls []GridLine) int {
	sum := 0

	for _, l := range ls {
		sum += l.Distance()
	}

	return sum
}

func (w Wire) FindCrossingPoints(w2 Wire) []Intersection {
	crossings := []Intersection{}
	for w1idx, w1Line := range w {
		for w2idx, w2Line := range w2 {
			if crossingPoint := w1Line.CrossesLine(w2Line); crossingPoint != nil {
				firstLength := WireLength(w[:w1idx])
				secondLength := WireLength(w2[:w2idx])
				totalLength := firstLength + secondLength
				crossings = append(crossings, Intersection{
					Point:      *crossingPoint,
					wireLength: totalLength,
				})
			}
		}
	}

	return crossings
}

func (l GridLine) CrossesLine(l2 GridLine) *Point {
	// Two lines going in the same direction can't cross each other (in this task)
	if l.LineType == l2.LineType {
		return nil
	}

	var p *Point

	var (
		hLine GridLine
		vLine GridLine
	)
	if l.LineType == Horizontal {
		hLine = l
		vLine = l2
	} else {
		hLine = l2
		vLine = l
	}

	crossesY := (hLine.Start.Y < vLine.Start.Y && hLine.Start.Y > vLine.End.Y ||
		hLine.Start.Y > vLine.Start.Y && hLine.Start.Y < vLine.End.Y)
	crossesX := (vLine.Start.X < hLine.Start.X && vLine.Start.X > hLine.End.X ||
		vLine.Start.X > hLine.Start.X && vLine.Start.X < hLine.End.X)
	if crossesX && crossesY {
		p = &Point{
			X: vLine.Start.X,
			Y: hLine.Start.Y,
		}
		fmt.Printf("the line %v crosses the line %v in both axis in the point %v\n", hLine, vLine, p)
	}

	return p
}

func CreateWire(wiring []string) Wire {
	var wire Wire
	start := Point{
		X: 0,
		Y: 0,
	}

	for _, part := range wiring {
		direction := string(part[0])
		steps, err := strconv.Atoi(string(part[1:]))
		util.CheckErr(err)

		var (
			end      Point
			lineType LineType
		)

		switch direction {
		case "U":
			end = Point{
				X: start.X,
				Y: start.Y + steps,
			}
			lineType = Vertical
		case "D":
			end = Point{
				X: start.X,
				Y: start.Y - steps,
			}
			lineType = Vertical
		case "R":
			end = Point{
				X: start.X + steps,
				Y: start.Y,
			}
			lineType = Horizontal
		case "L":
			end = Point{
				X: start.X - steps,
				Y: start.Y,
			}
			lineType = Horizontal
		default:
			panic(fmt.Sprintf("unrecognized direction: %v", direction))
		}
		line := GridLine{
			Start:    start,
			End:      end,
			LineType: lineType,
		}
		wire = append(wire, line)
		start = end
	}

	return wire
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
