package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestMergeSorted(t *testing.T) {
	testCases := []struct {
		name       string
		a, b, want []int
	}{
		{
			name: "merge two sorted arrays",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "left array finishes first",
			a:    []int{1, 2},
			b:    []int{3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "right array finishes first",
			a:    []int{1, 3, 5},
			b:    []int{2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicates are preserved",
			a:    []int{1, 2, 2},
			b:    []int{2, 3},
			want: []int{1, 2, 2, 2, 3},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := mergeSorted(tC.a, tC.b)
			assert.Equal(t, tC.want, got)
		})
	}
}
