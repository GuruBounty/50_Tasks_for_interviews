package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestLongestUniqueSubstringLen(t *testing.T) {
	testCases := []struct {
		desc, s string
		want    int
	}{
		{"\"\" => 0", "", 0},
		{"\"a\" => 1", "a", 1},
		{"\"ab\" => 2", "ab", 2},
		{"\"abb\" => 2", "abb", 2},
		{"\"abba\" => 2", "abba", 2},
		{"\"привет🙂мир🙂\" => 8", "привет🙂мир🙂", 8},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := longestUniqueSubstringLen(tC.s)
			assert.Equal(t, tC.want, got)
		})
	}
}
