package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestDedupPreserveOrder(t *testing.T) {
	testCases := []struct {
		desc    string
		s, want []string
	}{
		{desc: "{ab cd ef} => {ab cd ef}", s: []string{"ab", "cd", "ef"}, want: []string{"ab", "cd", "ef"}},
		{desc: "{ab cd ef ef} => {ab cd ef}", s: []string{"ab", "cd", "ef", "ef"}, want: []string{"ab", "cd", "ef"}},
		{desc: "{ab cd ef ef cd ab} => {ab cd ef}", s: []string{"ab", "cd", "ef", "ef", "cd", "ab"}, want: []string{"ab", "cd", "ef"}},
		{desc: "{ab ab dd hh} => {ab dd hh}", s: []string{"ab", "ab", "dd", "hh"}, want: []string{"ab", "dd", "hh"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := dedupPreserveOrder(tC.s)
			assert.Equal(t, tC.want, got)
		})
	}
}
