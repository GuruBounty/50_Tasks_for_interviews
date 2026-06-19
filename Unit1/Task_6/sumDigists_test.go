package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_sumDigists(t *testing.T) {
	testCases := []struct {
		desc      string
		num, want int
	}{
		{"1 => 1", 1, 1},
		{"10 => 1", 10, 1},
		{"12 => 3", 12, 3},
		{"123 => 6", 123, 6},
		{"1234 => 10", 1234, 10},
		{"-1 => 1", -1, 1},
		{"-10 => 1", -10, 1},
		{"-12 => 3", -12, 3},
		{"-123 => 6", -123, 6},
		{"-1234 => 10", -1234, 10},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := sumDigits(tC.num)
			assert.Equal(t, tC.want, result)
		})
	}
}
