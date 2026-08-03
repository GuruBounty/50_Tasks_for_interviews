package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestFindIndex(t *testing.T) {
	testCases := []struct {
		desc    string
		a       []int
		x, want int
	}{
		{"[1,2,3], 1 => 0", []int{1, 2, 3}, 1, 0},
		{"[1,2,3], 2 => 1", []int{1, 2, 3}, 2, 1},
		{"[1,2,3], 3 => 2", []int{1, 2, 3}, 3, 2},
		{"[-5 -1 0 3 8 12 20], 8 => 4", []int{-5, -1, 0, 3, 8, 12, 20}, 8, 4},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findIndex(tC.a, tC.x)
			assert.Equal(t, tC.want, got)
		})
	}
}
