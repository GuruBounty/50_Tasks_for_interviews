package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func Test_FahrenheitToCelsius(t *testing.T) {
	testCases := []struct {
		desc string
		f    any
		want float64
		err  error
	}{
		{"32 => 0, nil", 32.0, 0, nil},
		{"50 => 10, nil", 50.0, 10, nil},
		{"abc => 0, ivalid type", "abc", 0, fmt.Errorf("invalid type")},
		{"32 => 0, ivalid type", 32, 0, fmt.Errorf("invalid type")},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := FahrenheitToCelsius(tC.f)
			assert.Equal(t, tC.err, err)
			assert.Equal(t, tC.want, result)
		})
	}
}
