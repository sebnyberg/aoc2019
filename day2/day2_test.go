package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunIntegerMachine(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			"first",
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			res := RunIntegerMachine(tt.input)
			fmt.Println(res)
			require.Equal(t, tt.output, res)
		})
	}
}
