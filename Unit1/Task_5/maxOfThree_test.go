package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_maxOfThree(t *testing.T) {
	testCases := []struct {
		desc                   string
		num1, num2, num3, want int
	}{
		{"3,7,5 => 7", 3, 7, 5, 7},
		{"1,2,3 => 3", 1, 2, 3, 3},
		{"4,2,1 => 4", 4, 2, 1, 4},
		{"5,5,4 => 5", 5, 5, 4, 5},
		{"5,5,5 => 5", 5, 5, 5, 5},
		{"-5,-10,-50 => -5", -5, -10, -50, -5},
		{"-5,10,50 => 50", -5, 10, 50, 50},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := maxOfThree(tC.num1, tC.num2, tC.num3)
			assert.Equal(t, tC.want, result)
		})
	}
}
