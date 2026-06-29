package main

import (
	"sync"
	"testing"

	"github.com/alecthomas/assert"
)

func Test_mapper(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{desc: "single value", nums: []int{1}, want: 1},
		{desc: "two values", nums: []int{1, 2}, want: 5},
		{desc: "three values", nums: []int{1, 2, 3}, want: 14},
		{desc: "three values", nums: []int{1, 2, 3, 4}, want: 30},
		{desc: "three values", nums: []int{1, 2, 3, 4, 5}, want: 55},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			//defer wg.Wait()
			chOUt := make(chan int)
			//defer close(chOUt)
			go mapper(tC.nums, chOUt, &wg)

			result := <-chOUt
			wg.Wait()
			assert.Equal(t, tC.want, result)
		})
	}
}

func Test_reduce(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{desc: "single value", nums: []int{1}, want: 1},
		{desc: "single value", nums: []int{1, 2}, want: 3},
		{desc: "single value", nums: []int{1, 2, 3}, want: 6},
		{desc: "single value", nums: []int{1, 2, 3, 4}, want: 10},
		{desc: "single value", nums: []int{1, 2, 3, 4, 5}, want: 15},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			chIn := make(chan int, len(tC.nums))
			chOut := make(chan int)

			go reduce(chIn, chOut, &wg)
			for _, v := range tC.nums {
				chIn <- v
			}
			close(chIn)

			result := <-chOut
			wg.Wait()

			assert.Equal(t, tC.want, result)
		})
	}
}
