package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		desc       string
		a, b, want []int
	}{
		{"[1] [1] => [1]", []int{1}, []int{1}, []int{1}},
		{"[1] [1, 2] => [1]", []int{1}, []int{1, 2}, []int{1}},
		{"[1, 2] [1] => [1]", []int{1, 2}, []int{1}, []int{1}},
		{"[1, 2, 3] [1, 2, 4] => [1, 2]", []int{1, 2, 3}, []int{1, 2, 4}, []int{1, 2}},
		{"[1, -10, 3, 2] [1, 2, -10] => [-10, 1, 2]", []int{1, -10, 3, 2}, []int{1, 2, -10}, []int{-10, 1, 2}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := intersect(tC.a, tC.b)
			assert.Equal(t, tC.want, got)
		})
	}
}
