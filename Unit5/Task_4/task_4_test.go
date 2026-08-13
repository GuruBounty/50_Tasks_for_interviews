package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		desc    string
		a, want [][2]int
	}{
		{"[[1 3], [2 6] [8 10] [10 12] => [[1, 6], [8, 12]]", [][2]int{{1, 3}, {2, 6}, {8, 10}, {10, 12}}, [][2]int{{1, 6}, {8, 12}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := mergeIntervals(tC.a)
			assert.Equal(t, tC.want, got)
		})
	}
}
