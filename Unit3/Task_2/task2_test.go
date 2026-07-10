package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsAnagramUnicodeCaseInsensitive(t *testing.T) {
	testCases := []struct {
		desc, s, t string
		want       bool
	}{
		{"a, a => true", "a", "a", true},
		{"a👌, a👌 => true", "a👌", "a👌", true},
		{"a, ab => false", "a", "ab", false},
		{"a🤔, a😒 => false", "a🤔", "a😒", false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isAnagramUnicodeCaseInsensitive(tC.s, tC.t)
			assert.Equal(t, tC.want, got)
		})
	}
}
