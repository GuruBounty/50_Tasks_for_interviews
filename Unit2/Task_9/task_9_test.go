package main

import (
	"testing"
	"time"
)

// TestTryRecv_EmptyChannel tests receiving from an empty channel
func TestTryRecv_EmptyChannel(t *testing.T) {
	ch := make(chan int, 5)
	val, ok := TryRecv(ch)

	if ok {
		t.Errorf("Expected ok=false for empty channel, got true")
	}
	if val != 0 {
		t.Errorf("Expected val=0 for empty channel, got %d", val)
	}
}

// TestTryRecv_NonEmptyChannel tests receiving from a non-empty channel
func TestTryRecv_NonEmptyChannel(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 42

	val, ok := TryRecv(ch)

	if !ok {
		t.Errorf("Expected ok=true, got false")
	}
	if val != 42 {
		t.Errorf("Expected val=42, got %d", val)
	}
}

// TestTryRecv_MultipleValues tests receiving from a channel with multiple values
func TestTryRecv_MultipleValues(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 10
	ch <- 20
	ch <- 30

	val1, ok1 := TryRecv(ch)
	val2, ok2 := TryRecv(ch)
	val3, ok3 := TryRecv(ch)
	val4, ok4 := TryRecv(ch)

	if !ok1 || val1 != 10 {
		t.Errorf("First recv: expected (10, true), got (%d, %v)", val1, ok1)
	}
	if !ok2 || val2 != 20 {
		t.Errorf("Second recv: expected (20, true), got (%d, %v)", val2, ok2)
	}
	if !ok3 || val3 != 30 {
		t.Errorf("Third recv: expected (30, true), got (%d, %v)", val3, ok3)
	}
	if ok4 {
		t.Errorf("Fourth recv: expected (0, false), got (%d, %v)", val4, ok4)
	}
}

// TestTryRecv_UnbufferedChannel tests receiving from an unbuffered channel
func TestTryRecv_UnbufferedChannel(t *testing.T) {
	ch := make(chan int)

	// Should fail immediately on empty unbuffered channel
	val, ok := TryRecv(ch)
	if ok {
		t.Errorf("Expected ok=false for empty unbuffered channel, got true")
	}
	if val != 0 {
		t.Errorf("Expected val=0, got %d", val)
	}
}

// TestTryRecv_WithGoroutine tests receiving when a goroutine is sending
func TestTryRecv_WithGoroutine(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch <- 100
	}()

	// First try should fail
	_, ok1 := TryRecv(ch)
	if ok1 {
		t.Errorf("Expected first recv to fail, got ok=true")
	}

	// Wait for the goroutine to send, then try again
	time.Sleep(50 * time.Millisecond)
	val2, ok2 := TryRecv(ch)
	if !ok2 || val2 != 100 {
		t.Errorf("Expected second recv to succeed with 100, got (%d, %v)", val2, ok2)
	}
}

// TestTrySend_BufferedChannelWithCapacity tests sending to a buffered channel with available space
func TestTrySend_BufferedChannelWithCapacity(t *testing.T) {
	ch := make(chan int, 5)

	ok := TrySend(ch, 42)
	if !ok {
		t.Errorf("Expected ok=true for sending to empty buffered channel, got false")
	}

	val, recvOk := TryRecv(ch)
	if !recvOk || val != 42 {
		t.Errorf("Expected to receive 42, got (%d, %v)", val, recvOk)
	}
}

// TestTrySend_BufferedChannelFull tests sending to a full buffered channel
func TestTrySend_BufferedChannelFull(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// Channel is now full
	ok := TrySend(ch, 3)
	if ok {
		t.Errorf("Expected ok=false for full channel, got true")
	}

	// Verify original values are still in channel
	val1, _ := TryRecv(ch)
	val2, _ := TryRecv(ch)
	if val1 != 1 || val2 != 2 {
		t.Errorf("Expected to receive 1 and 2, got %d and %d", val1, val2)
	}
}

// TestTrySend_UnbufferedChannel tests sending to an unbuffered channel
func TestTrySend_UnbufferedChannel(t *testing.T) {
	ch := make(chan int)

	// Should fail immediately on unbuffered channel without receiver
	ok := TrySend(ch, 42)
	if ok {
		t.Errorf("Expected ok=false for unbuffered channel without receiver, got true")
	}
}

// TestTrySend_MultipleValues tests sending multiple values
func TestTrySend_MultipleValues(t *testing.T) {
	ch := make(chan int, 3)

	ok1 := TrySend(ch, 10)
	ok2 := TrySend(ch, 20)
	ok3 := TrySend(ch, 30)
	ok4 := TrySend(ch, 40) // Should fail - channel is full

	if !ok1 || !ok2 || !ok3 {
		t.Errorf("Expected first 3 sends to succeed: ok1=%v, ok2=%v, ok3=%v", ok1, ok2, ok3)
	}
	if ok4 {
		t.Errorf("Expected 4th send to fail, got ok=true")
	}

	// Verify all sent values
	val1, _ := TryRecv(ch)
	val2, _ := TryRecv(ch)
	val3, _ := TryRecv(ch)

	if val1 != 10 || val2 != 20 || val3 != 30 {
		t.Errorf("Expected 10, 20, 30, got %d, %d, %d", val1, val2, val3)
	}
}

// TestTrySend_WithGoroutine tests sending when a goroutine is receiving
func TestTrySend_WithGoroutine(t *testing.T) {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Millisecond)
		val := <-ch
		if val != 100 {
			t.Errorf("Expected to receive 100, got %d", val)
		}
		done <- true
	}()

	// First try should fail (no receiver ready immediately)
	ok := TrySend(ch, 100)
	if ok {
		t.Errorf("Expected first send to fail, got ok=true")
	}

	// Wait for goroutine to finish receiving
	<-done
}

// TestTrySend_NegativeValue tests sending a negative value
func TestTrySend_NegativeValue(t *testing.T) {
	ch := make(chan int, 1)

	ok := TrySend(ch, -5)
	if !ok {
		t.Errorf("Expected ok=true for sending negative value, got false")
	}

	val, recvOk := TryRecv(ch)
	if !recvOk || val != -5 {
		t.Errorf("Expected to receive -5, got (%d, %v)", val, recvOk)
	}
}

// TestTrySend_ZeroValue tests sending a zero value
func TestTrySend_ZeroValue(t *testing.T) {
	ch := make(chan int, 1)

	ok := TrySend(ch, 0)
	if !ok {
		t.Errorf("Expected ok=true for sending zero value, got false")
	}

	val, recvOk := TryRecv(ch)
	if !recvOk || val != 0 {
		t.Errorf("Expected to receive 0, got (%d, %v)", val, recvOk)
	}
}
