package fib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fib(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
	}

	for _, c := range cases {
		assert.Equal(t, c.output, Fib(c.input), fmt.Sprintf("input:%d, output:%d", c.input, c.output))
	}
}
