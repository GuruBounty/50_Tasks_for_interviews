package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestRemoveElemtInSlice(t *testing.T) {
	testCases := []struct {
		desc, want string
		s          []string
		idx        int
	}{
		{desc: "{ab cd ef}, 2 => ab cd", s: []string{"ab", "cd", "ef"}, idx: 2, want: "ab cd"},
		{desc: "{ab cd ef}, 1 => ab ef", s: []string{"ab", "cd", "ef"}, idx: 1, want: "ab ef"},
		{desc: "{ab cd ef}, 0 => cd ef", s: []string{"ab", "cd", "ef"}, idx: 0, want: "cd ef"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := RemoveElemtInSlice(tC.s, tC.idx)
			assert.Equal(t, tC.want, got)
		})
	}
}
