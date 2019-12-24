package day3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sebnyberg/aoc2019/day3"
	"github.com/sebnyberg/aoc2019/util"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestDay3(t *testing.T) {
	fileWires := util.ReadFile("day3_input")
	first := strings.Split(fileWires[0], ",")
	second := strings.Split(fileWires[1], ",")

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
				first,
				second,
			},
			expected: 4981,
		},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("test_%v", idx), func(t *testing.T) {
			firstWire := day3.CreateWire(tc.input[0])
			secondWire := day3.CreateWire(tc.input[1])

			fmt.Printf("wire two: %v\n", firstWire)
			fmt.Printf("wire one: %v\n", secondWire)

			minDistance := 1000000000

			startingPoint := day3.Point{
				X: 0,
				Y: 0,
			}
			crossingPoints := firstWire.FindCrossingPoints(secondWire)
			require.Greater(t, len(crossingPoints), 1)
			for _, crossing := range firstWire.FindCrossingPoints(secondWire) {
				distance := startingPoint.DistanceTo(crossing.Point)
				if distance < minDistance {
					fmt.Printf("new min distance for points %v and %v, distance: %v\n", startingPoint, crossing, distance)
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
