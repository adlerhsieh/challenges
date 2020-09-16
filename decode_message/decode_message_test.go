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
			input:    "",
			expected: 1,
		},
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
			input:    "1234",
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
			// 2,1,1,1
			// 21,1,1
			// 21,11
			// 2,11,1
			// 2,1,11
		},
		{
			input:    "11311",
			expected: 9,
			// 1,1,3,1,1
			// 11,3,1,1
			// 11,3,11
			// 11,31,1
			// 1,13,1,1
			// 1,13,11
			// 1,1,31,1
			// 11,31,1
			// 1,1,3,11
		},
	}
	for _, c := range cases {
		fmt.Println("=== input:", c.input)
		actual := NumWays(c.input)
		if c.expected != actual {
			t.Fatalf(fmt.Sprintf("input: %s, expected: %d, actual: %d", c.input, c.expected, actual))
		}
	}
}
