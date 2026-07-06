package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_fanIn(t *testing.T) {
	type testCase struct {
		desc   string
		inputA []int
		inputB []int
		inputC []int
	}

	testCases := []testCase{
		{
			desc:   "three non-empty channels",
			inputA: []int{1, 2, 3},
			inputB: []int{4},
			inputC: []int{5, 6},
		},
		{
			desc:   "empty first channel",
			inputA: []int{},
			inputB: []int{7, 8},
			inputC: []int{9},
		},
		{
			desc:   "all channels empty",
			inputA: []int{},
			inputB: []int{},
			inputC: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			chanA := make(chan int)
			chanB := make(chan int)
			chanC := make(chan int)

			out := fanIn(chanA, chanB, chanC)

			go func() {
				for _, v := range tc.inputA {
					chanA <- v
				}
				close(chanA)
			}()

			go func() {
				for _, v := range tc.inputB {
					chanB <- v
				}
				close(chanB)
			}()

			go func() {
				for _, v := range tc.inputC {
					chanC <- v
				}
				close(chanC)
			}()

			got := make([]int, 0)
			for v := range out {
				got = append(got, v)
			}

			expected := append(append(append([]int{}, tc.inputA...), tc.inputB...), tc.inputC...)
			sort.Ints(got)
			sort.Ints(expected)

			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("fanIn output = %v, want %v", got, expected)
			}
		})
	}
}
