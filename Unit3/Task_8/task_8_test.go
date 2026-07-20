package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestAtoi32(t *testing.T) {
	testCases := []struct {
		desc, s string
		want    int32
		ok      bool
	}{
		{"\"1\" => 1,  true", "1", 1, true},
		{"\"12😊\" => 0, false", "12😊", 0, false},
		{"\"2147483647\" => 2147483647, true", "2147483647", 2147483647, true},
		{"\"-1\" => -1, true", "-1", -1, true},
		{"\"-1❤️\" => 0, false", "-1❤️", 0, false},
		{"\"-2147483648\" => -2147483648, true", "-2147483648", -2147483648, true},
		{"\"2147483648\" => 0, false", "2147483648", 0, false},
		{"\"-2147483649\" => 0, false", "-2147483649", 0, false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, res := atoi32(tC.s)
			assert.Equal(t, tC.want, got)
			assert.Equal(t, tC.ok, res)
		})
	}
}
