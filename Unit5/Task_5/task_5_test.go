package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestFindMinElement(t *testing.T) {
	testCases := []struct {
		desc string
		a    []int
		want int
	}{
		{"[3 4 5 1 2] => 1", []int{3, 4, 5, 1, 2}, 1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findMinElement(tC.a)
			assert.Equal(t, tC.want, got)
		})
	}
}
