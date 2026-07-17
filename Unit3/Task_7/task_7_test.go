package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestNormalizeSpaces(t *testing.T) {
	testCases := []struct {
		desc, s, want string
	}{
		{"a => a", "a", "a"},
		{"Fh😨 \n1x => Fh😨 1x", "Fh😨 \n1x", "Fh😨 1x"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := normalizeSpaces(tC.s)
			assert.Equal(t, tC.want, got)
		})
	}
}
