package decode_message

import (
	"fmt"
	"testing"
)

func Test_NumWays(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "1",
			expected: 1,
		},
		{
			input:    "2",
			expected: 1,
		},
		{
			input:    "12",
			expected: 2,
		},
		{
			input:    "123",
			expected: 3,
		},
		{
			input:    "243",
			expected: 2,
		},
		{
			input:    "2432",
			expected: 2,
		},
		{
			input:    "2111",
			expected: 5,
		},
	}
	for _, c := range cases {
		actual := NumWays(c.input)
		if c.expected != actual {
			t.Fatalf(fmt.Sprintf("input: %s, expected: %d, actual: %d", c.input, c.expected, actual))
		}
	}
}
