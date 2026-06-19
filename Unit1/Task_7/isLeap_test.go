package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_isLeap(t *testing.T) {
	testCases := []struct {
		desc string
		year int
		want bool
	}{
		{"2024 => true", 2024, true},
		{"2025 => false", 2025, false},
		{"2026 => false", 2026, false},
		{"2027 => false", 2027, false},
		{"2028 => true", 2028, true},
		{"2000 => true", 2000, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := isLeap(tC.year)
			assert.Equal(t, tC.want, result)
		})
	}
}
