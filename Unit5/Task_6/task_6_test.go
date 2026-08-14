package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestQuickSelect(t *testing.T) {
	testCases := []struct {
		desc    string
		a       []int
		k, want int
	}{
		{"[3 2 1 5 6 4], 2 => 5", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"[3 2 3 1 2 4 5 5 6], 4 => 4", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := quickSelect(tC.a, tC.k)
			assert.Equal(t, tC.want, got)
		})
	}
}
