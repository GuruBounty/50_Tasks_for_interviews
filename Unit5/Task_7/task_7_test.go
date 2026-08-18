package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestBraces(t *testing.T) {
	testCases := []struct {
		desc, line, want string
	}{
		{"{} => YES", "{}", "YES"},
		{"() => YES", "()", "YES"},
		{"[] => YES", "[]", "YES"},
		{") => NO", ")", "NO"},
		{"] => NO", "]", "NO"},
		{"} => NO", "}", "NO"},
		{"()} => NO", "()}", "NO"},
		{"{[()()]} => YES", "{[()()]}", "YES"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := braces(tC.line)
			assert.Equal(t, tC.want, got)
		})
	}
}
