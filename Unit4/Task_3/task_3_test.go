package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		desc    string
		a, want []int
		k       int
	}{
		{desc: "{1,1}, 1 => {1}", a: []int{1, 1}, k: 1, want: []int{1}},
		{desc: "{}, 1 => {}", a: []int{}, k: 1, want: []int{}},
		{desc: "{1,1}, 0 => {}", a: []int{1, 1}, k: 0, want: []int{}},
		{desc: "{2,2}, 1 => {2}", a: []int{2, 2}, k: 1, want: []int{2}},
		{desc: "{2,2,1,1}, 2 => {1,2}", a: []int{2, 2, 1, 1}, k: 2, want: []int{1, 2}},
		{desc: "{1,1,1,2,2,3}, 2 => {1,2}", a: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := topKFrequent(tC.a, tC.k)
			assert.Equal(t, tC.want, got)
		})
	}
}
