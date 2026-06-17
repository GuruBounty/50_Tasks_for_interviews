package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_countVowels(t *testing.T) {
	testCases := []struct {
		desc, s string
		want    int
	}{
		{"a => 1", "a", 1},
		{"ab => 1", "ab", 1},
		{"aа => 2", "aа", 2},
		{"aeiou => 5", "aeiou", 5},
		{"AeIouаеЁиОуЫэЮя => 15", "AeIouаеЁиОуЫэЮя", 15},
		{"Hello => 2", "Hello", 2},
		{"\"\" => 0", "", 0},
		{"TRmnwcxzPLJG => 0", "TRmnwcxzPLJG", 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := countVowels(tC.s)
			assert.Equal(t, tC.want, result)
		})
	}
}
