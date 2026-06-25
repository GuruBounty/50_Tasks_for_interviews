package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_generator(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{desc: "single value", nums: []int{42}, want: []int{42}},
		{desc: "1, 2, 3", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)

			out := make(chan int, len(tc.nums))
			go generator(tc.nums, out, &wg)

			wg.Wait()

			var got []int
			for v := range out {
				got = append(got, v)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("generator() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_multiplier(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{desc: "single value", nums: []int{1}, want: []int{2}},
		{desc: "multiple values", nums: []int{1, 2, 3}, want: []int{2, 4, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)

			in := make(chan int, len(tc.nums))
			out := make(chan int, len(tc.nums))

			go multiplier(in, out, &wg)

			for _, num := range tc.nums {
				in <- num
			}
			close(in)

			wg.Wait()

			var got []int
			for v := range out {
				got = append(got, v)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("multiplier() = %v, want %v", got, tc.want)
			}
		})
	}
}
