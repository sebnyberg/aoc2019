package day3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sebnyberg/aoc2019/day3"
	"github.com/sebnyberg/aoc2019/util"
	"gotest.tools/assert"
)

func TestDay3(t *testing.T) {
	fileWires := util.ReadFile("day3_input")

	firstWire := strings.Split(fileWires[0], ",")
	secondWire := strings.Split(fileWires[1], ",")

	tcs := []struct {
		input    [][]string
		expected int
	}{
		{
			input: [][]string{
				[]string{"R8", "U5", "L5", "D3"},
				[]string{"U7", "R6", "D4", "L4"},
			},
			expected: 6,
		},
		{
			input: [][]string{
				[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
				[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			},
			expected: 159,
		},
		{
			input: [][]string{
				[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
				[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			},
			expected: 135,
		},
		{
			input: [][]string{
				firstWire,
				secondWire,
			},
			expected: 10,
		},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("test_%v", idx), func(t *testing.T) {
			grid := day3.NewGrid()

			port := grid.GetPort()
			grid.PutWire(tc.input[0])
			grid.PutWire(tc.input[1])

			crossings := grid.GetCrossings()
			fmt.Println(crossings)

			minDistance := 1024 * 2

			for _, crossing := range crossings {
				distance := port.DistanceTo(crossing)
				if distance < minDistance {
					fmt.Printf("new min distance for points %v and %v, distance: %v\n", port, crossing, distance)
					minDistance = distance
				}
			}

			assert.Equal(t, tc.expected, minDistance)
		})
	}
}

// func TestFile(t *testing.T) {
// 	fileContent := util.ReadFile("day3_input")
// 	firstInput := strings.Split(fileContent[0], ",")
// 	secondInput := strings.Split(fileContent[1], ",")
// }
