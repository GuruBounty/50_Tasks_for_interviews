package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestRleCompressRunes(t *testing.T) {
	testCases := []struct {
		desc, s, want string
	}{
		{"\"\" => \"\"", "", ""},
		{"\"a\" => a1", "a", "a1"},
		{"\"aa\" => a2", "aa", "a2"},
		{"\"aab\" => a2b1", "aab", "a2b1"},
		{"\"aaabb\" => a3b2", "aaabb", "a3b2"},
		{"\"abbccaa\" => a1b2c2a2", "abbccaa", "a1b2c2a2"},
		{"\"🙂🙂🙂aa\" => 🙂3a2", "🙂🙂🙂aa", "🙂3a2"},
		{"\"абббвв\" => а1б3в2", "абббвв", "а1б3в2"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := rleCompressRunes(tC.s)
			assert.Equal(t, tC.want, got)
		})
	}
}
