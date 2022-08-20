package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: -1,
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:   1,
			expected: 0,
		},
	}

	for _, tt := range cases {
		fmt.Println("input:", tt.nums)
		fmt.Println("target:", tt.target)
		fmt.Println("expected:", tt.expected)
		result := BinarySearch(tt.nums, tt.target)
		if result != tt.expected {
			t.Fatalf("test failed. expected: %d, actual: %d, input: %v, target: %d", tt.expected, result, tt.nums, tt.target)
		}
	}
}
