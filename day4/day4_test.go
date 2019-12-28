package day4_test

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/aoc2019/day4"
	"github.com/stretchr/testify/require"
)

func Test_Day4Part2(t *testing.T) {
	tcs := []struct {
		in       int
		expected bool
	}{
		{111111, false},
		{223450, false},
		{123789, false},
		{112233, true},
		{123444, false},
		{111122, true},
		{111222, false},
		{112222, true},
	}

	require.Equal(t, 1, day4.Part2(193651, 649729))

	for _, tc := range tcs {
		t.Run(fmt.Sprint("test_", tc.in), func(t *testing.T) {
			require.Equal(t, tc.expected, day4.CheckNumber(tc.in))
		})
	}
}
