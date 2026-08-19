package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestSumSubSlice(t *testing.T) {
	testCases := []struct {
		desc string
		a    []int
		K    int
		want string
	}{
		{"[1,1], 2 => YES", []int{1, 1}, 2, "YES"},
		{"[1,1], 3 => NO", []int{1, 1}, 3, "NO"},
		{"[1 2 3 -2 5], 3 => YES", []int{1, 2, 3, -2, 5}, 3, "YES"},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := sumSubSlice(tC.a, tC.K)
			assert.Equal(t, tC.want, got)
		})
	}
}
