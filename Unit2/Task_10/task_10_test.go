package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_Solve(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{name: "single task", nums: []int{3}, k: 1, want: 9},
		{name: "multiple tasks", nums: []int{1, 2, 3, 4}, k: 2, want: 30},
		{name: "empty input", nums: nil, k: 3, want: 0},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := solve(tC.nums, tC.k)
			assert.Equal(t, got, tC.want)
			// if got := solve(tt.nums, tt.k); got != tt.want {
			// 	t.Fatalf("solve(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			// }
		})
	}
}
