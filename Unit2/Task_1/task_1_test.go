package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_parallelSearch(t *testing.T) {
	testCases := []struct {
		desc     string
		arr      []int
		target   int
		expected int
	}{
		{
			desc:     "nil slice returns -1",
			arr:      nil,
			target:   5,
			expected: -1,
		},
		{
			desc:     "empty slice returns -1",
			arr:      []int{},
			target:   0,
			expected: -1,
		},
		{
			desc:     "target in left half",
			arr:      []int{3, 1, 4, 2, 5},
			target:   1,
			expected: 1,
		},
		{
			desc:     "target in right half",
			arr:      []int{3, 1, 4, 2, 5},
			target:   2,
			expected: 3,
		},
		{
			desc:     "target not present",
			arr:      []int{7, 8, 9, 10},
			target:   5,
			expected: -1,
		},
		{
			desc:     "duplicate values returns first index",
			arr:      []int{5, 5, 5, 5},
			target:   5,
			expected: 0,
		},
		{
			desc:     "odd length target at boundary",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := parallelSearch(tc.arr, tc.target)
			assert.Equal(t, tc.expected, result)

		})
	}
}
