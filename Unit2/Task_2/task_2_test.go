package main

import (
	"sync"
	"testing"
	"time"
)

func Test_worker(t *testing.T) {
	testCases := []struct {
		desc     string
		id       int
		workTime time.Duration
	}{
		{
			desc:     "single worker completes",
			id:       1,
			workTime: 20 * time.Millisecond,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)

			done := make(chan struct{})
			start := time.Now()

			go func() {
				worker(tc.id, tc.workTime, &wg)
				close(done)
			}()

			waitDone := make(chan struct{})
			go func() {
				wg.Wait()
				close(waitDone)
			}()

			select {
			case <-done:
			case <-time.After(time.Second):
				t.Fatal("worker did not finish in time")
			}

			select {
			case <-waitDone:
			case <-time.After(time.Second):
				t.Fatal("wait group was not released")
			}

			elapsed := time.Since(start)
			if elapsed < tc.workTime {
				t.Fatalf("worker finished too early: got %v, want at least %v", elapsed, tc.workTime)
			}
		})
	}
}
