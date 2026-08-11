package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestFoundTwoElemtsIfSumEqualTarget(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   string
	}{
		{"[], 10 => NO", []int{}, 10, "NO"},
		{"[1,1], 2 => YES", []int{1, 1}, 2, "YES"},
		{"[1,2], 3 => YES", []int{1, 2}, 3, "YES"},
		{"[2 7 11 15 3], 9 => YES", []int{2, 7, 11, 15, 3}, 9, "YES"},
		{"[2 7 11 15 3], 19 => YES", []int{2, 7, 11, 15, 3}, 19, "NO"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := FoundTwoElemtsIfSumEqualTarget(tC.nums, tC.target)
			assert.Equal(t, tC.want, got)
		})
	}
}
