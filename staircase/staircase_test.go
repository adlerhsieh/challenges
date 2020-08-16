package staircase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_St(t *testing.T) {
	cases := []struct {
		n          int
		approaches [][]int
	}{
		{
			n: 1,
			approaches: [][]int{
				[]int{1},
			},
		},
		{
			n: 2,
			approaches: [][]int{
				[]int{1, 1},
				[]int{2},
			},
		},
		{
			n: 3,
			approaches: [][]int{
				[]int{1, 1, 1},
				[]int{2, 1},
				[]int{1, 2},
			},
		},
		{
			n: 4,
			approaches: [][]int{
				[]int{1, 1, 1, 1},
				[]int{1, 2, 1},
				[]int{1, 1, 2},
				[]int{2, 2},
			},
		},
	}

	for _, c := range cases {
		result := St(c.n)
		assert.Equal(t, len(c.approaches), len(result), fmt.Sprintf("n = %d", c.n))
	}
}
