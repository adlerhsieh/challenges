package staircase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_St(t *testing.T) {
	cases := []struct {
		input  int
		ways   []int
		output int
	}{
		{
			input:  2,
			ways:   []int{1, 2},
			output: 2,
		},
		{
			input:  3,
			ways:   []int{1, 2, 3},
			output: 4,
		},
		{
			input:  3,
			ways:   []int{1},
			output: 1,
		},
		{
			input:  4,
			ways:   []int{3, 5},
			output: 1,
		},
		{
			input:  4,
			ways:   []int{1, 2, 5},
			output: 5,
		},
		{
			input:  4,
			ways:   []int{1, 2, 3},
			output: 7,
		},
	}

	for _, c := range cases {
		result := St(c.input, c.ways)
		assert.Equal(t, c.output, result, fmt.Sprintf("n = %d, ways = %v", c.input, c.ways))
	}
}
