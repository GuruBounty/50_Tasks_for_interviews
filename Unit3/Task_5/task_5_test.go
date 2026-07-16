package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestPerfixSuffixStatus(t *testing.T) {
	testCases := []struct {
		desc, s, p, want string
	}{
		{"p > s => NONE", "a", "😚a", "NONE"},
		{"p is empty => BOTH", "a", "", "BOTH"},
		{"p is prefix of s => PREFIX", "abcd", "ab", "PREFIX"},
		{"p is suffix of s => SUFFIX", "abcd", "cd", "SUFFIX"},
		{"p is prefix of s => PREFIX", "абвг", "аб", "PREFIX"},
		{"p and s => BOTH", "🤔ab🤔", "🤔", "BOTH"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := prefixSuffixStatus(tC.s, tC.p)
			assert.Equal(t, tC.want, got)
		})
	}
}
