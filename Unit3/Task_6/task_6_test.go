package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestParseCSVLine(t *testing.T) {
	testCases := []struct {
		desc, line string
		want       []string
	}{
		{"empty", "", nil},
		{"\"a,b\" => {a,b}", "\"a,b\"", []string{"a", "b"}},
		{"БюЗк,6,,😽Ж😴oнВ => {БюЗк, 6,,😽Ж😴oнВ,,}", "БюЗк,6,,😽Ж😴oнВ,", []string{"БюЗк", "6", "", "😽Ж😴oнВ", ""}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseCSVLine(tC.line)
			assert.Equal(t, tC.want, got)
		})
	}
}
