package main

import "testing"

func TestSloveProcessesKTasks(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	cnt, sum := slove(nums, 3, 3)
	if cnt != 3 {
		t.Fatalf("expected 3 tasks processed, got %d", cnt)
	}
	if sum != 14 {
		t.Fatalf("expected sum 14, got %d", sum)
	}
}

func TestSloveZeroCancel(t *testing.T) {
	nums := []int{10, 20, 30}
	cnt, sum := slove(nums, 2, 0)
	if cnt != 0 {
		t.Fatalf("expected 0 tasks processed, got %d", cnt)
	}
	if sum != 0 {
		t.Fatalf("expected sum 0, got %d", sum)
	}
}

func TestSloveHandlesMoreWorkersThanTasks(t *testing.T) {
	nums := []int{2, 4}
	cnt, sum := slove(nums, 10, 2)
	if cnt != 2 {
		t.Fatalf("expected 2 tasks processed, got %d", cnt)
	}
	if sum != 20 {
		t.Fatalf("expected sum 20, got %d", sum)
	}
}
