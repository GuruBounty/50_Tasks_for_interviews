package main

import (
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestSolveDoesNotDeadlock(t *testing.T) {
	nums := make([]int, 50)
	for i := range nums {
		nums[i] = i + 1
	}

	done := make(chan int64, 1)
	go func() {
		done <- solve(nums, 3)
	}()

	select {
	case got := <-done:
		want := int64(42925)
		if got != want {
			t.Fatalf("solve() = %d, want %d", got, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("solve() deadlocked")
	}
}

func Test_taskValue(t *testing.T) {
	testCases := []struct {
		desc string
		task int
		want int64
	}{
		{"0 => 0", 0, 0},
		{"2 => 4", 2, 4},
		{"50 => 2500", 50, 2500},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := taskValue(tC.task)
			assert.Equal(t, tC.want, result)
		})
	}
}
