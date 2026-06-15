package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_isNumPalindrome(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
		want bool
	}{
		{"-123 => false", -123, false},
		{"121 => true", 121, true},
		{"59 => false", 59, false},
		{"88 => true", 88, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := isNumPalindrome(tC.num)
			assert.Equal(t, result, tC.want)
		})
	}
}
