package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_saveDivide(t *testing.T) {
	testCases := []struct {
		desc             string
		num1, num2, want int
		err              error
	}{
		{"10, 2 => 5, nil", 10, 2, 5, nil},
		{"10, 0 => 0, Divided by zero", 10, 0, 0, ErrDividByZero},
		{"10, -2 => -5, nil", 10, -2, -5, nil},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := saveDivid(tC.num1, tC.num2)
			assert.Equal(t, tC.err, err)
			assert.Equal(t, tC.want, result)
		})
	}
}
