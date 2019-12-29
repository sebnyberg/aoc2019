package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Perm(t *testing.T) {
	tcs := []struct {
		in  []int
		out [][]int
	}{
		{
			in: []int{1, 2, 3},
			out: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 2, 1},
				[]int{3, 1, 2},
			},
		},
		{
			in: []int{1, 2},
			out: [][]int{
				[]int{1, 2},
				[]int{2, 1},
			},
		},
	}
	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("test_%d", idx), func(t *testing.T) {
			res := day7part2.GetAllPerms(tc.in)
			require.Equal(t, tc.out, res)
		})
	}
}
