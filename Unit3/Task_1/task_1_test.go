package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_reverseRunes(t *testing.T) {
	testCases := []struct {
		desc, line, want string
	}{
		{"a => a", "a", "a"},
		{"Hello, 世界 🌍 => 🌍 界世 ,olleH", "Hello, 世界 🌍", "🌍 界世 ,olleH"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := reverseRunes(tC.line)
			assert.Equal(t, tC.want, got)
		})
	}
}
