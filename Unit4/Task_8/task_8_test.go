package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestCounter(t *testing.T) {
	testCases := []struct {
		desc  string
		c     Counter
		delta int
		want  int
	}{
		{"Value = 1 Inc = 1 => Get = 2", Counter{Value: 1}, 1, 2},
		{"Value = 0 Inc 5 => Get = 5", Counter{Value: 0}, 5, 5},
		{"Value = 10 Inc -3 => Get = 7", Counter{Value: 10}, -3, 7},
		{"Zero delta doesn't change value", Counter{Value: 1}, 0, 1},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.c.Inc(tC.delta)
			got := tC.c.Get()
			assert.Equal(t, tC.want, got)
		})
	}
}
