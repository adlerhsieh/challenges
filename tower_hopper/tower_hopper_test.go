package towerhopper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHoppable(t *testing.T) {
	cases := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 0},
			output: false,
		},
		{
			input:  []int{1, 1},
			output: true,
		},
		{
			input:  []int{0, 3},
			output: false,
		},
		{
			input:  []int{1, 2, 0, 1},
			output: true,
		},
		{
			input:  []int{4, 2, 0, 0, 2, 0},
			output: true,
		},
		{
			input:  []int{1, 1, 1, 1, 1, 1, 5, 0, 0, 0, 0, 2, 0, 3},
			output: true,
		},
		{
			input:  []int{1, 1, 1, 1, 1, 1, 5, 0, 0, 0, 0, 2, 0, 0},
			output: false,
		},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.output, isHoppable(tt.input), fmt.Sprintf("input: %v", tt.input))
	}
}
