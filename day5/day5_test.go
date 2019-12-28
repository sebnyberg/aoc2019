package main

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/aoc2019/day4"
	"github.com/stretchr/testify/require"
)

func Test_Day5Part1(t *testing.T) {
	tcs := []struct {
		in          interface{}
		expectedOut interface{}
	}{
		{},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprint("test_", tc.in), func(t *testing.T) {
			require.Equal(t, tc.expected, day4.CheckNumber(tc.in))
		})
	}
}
