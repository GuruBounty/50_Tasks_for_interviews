package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestThreeSumEqualZero(t *testing.T) {
	testCases := []struct {
		desc string
		a    []int
		want string
	}{
		{"[] => NO", []int{}, "NO"},
		{"[1] => NO", []int{1}, "NO"},
		{"[2] => NO", []int{2}, "NO"},
		{"[1, 2] => NO", []int{1, 2}, "NO"},
		{"[1, 2, 3] => NO", []int{1, 2, 3}, "NO"},
		{"[1, 2, -3] => YES", []int{1, 2, -3}, "YES"},
		{"[-1 0 1 2 -1 -4] => YES", []int{-1, 0, 1, 2, -1, -4}, "YES"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ThreeSumEqualZero(tC.a)
			assert.Equal(t, tC.want, got)
		})
	}
}
