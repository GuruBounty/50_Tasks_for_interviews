package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestCanonicalKey(t *testing.T) {
	testCases := []struct {
		desc, word, want string
	}{
		{
			"khfo", "khfo", "fhko",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := canonicalKey(tC.word)
			assert.Equal(t, tC.want, got)

		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		desc  string
		words []string
		want  [][]string
	}{
		{"khfo", []string{"khfo"}, [][]string{{"khfo"}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := groupAnagrams(tC.words)
			assert.Equal(t, tC.want, got)
		})
	}
}
