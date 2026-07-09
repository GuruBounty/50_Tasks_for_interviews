package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func taskValue(x int) int64 {
	return int64(x) * int64(x)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, k int
	if _, err := fmt.Fscan(in, &m, &k); err != nil {
		return
	}
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &nums[i])
	}
	fmt.Print(solve(nums, k))
}

func solve(nums []int, k int) int64 {
	// we are given:
	// -nums: a list of tasks
	// k: maximum number of concurrently running goroutines

	// We need to:
	// 1. Process all tasks in separate goroutines
	// 2. Ensure that no more than k goroutines run simultaneously
	// 3. Correctly calculate the sum of the results

	//Key idea:
	// Use a buffered channel like a semaphore
	// sem <- struct{}{} - ocupy a slot
	// <-sem - free a slot
	// results sent to channel results
	if k <= 0 {
		return 0
	}

	sem := make(chan struct{}, k)

	//results is a channel for sending results
	results := make(chan int64, len(nums))

	var wg sync.WaitGroup

	// Run a summing goroutne which reads from results and adds the values to sum
	// use a sepatrate WaitGroup or another way to wait for it to finish
	var sum int64
	var sumWg sync.WaitGroup
	sumWg.Add(1)

	go func() {
		defer sumWg.Done()
		for value := range results {
			sum += value
		}
	}()

	// For each task:
	// - Acquire a slot in the semaphore (sem <- struct{}{})
	// - increse the wg
	// - Run a goroutine which:
	//   - calculate taskValue
	//   - send the result to results channel
	//   - free the slot in the semaphore(<-sem)
	//	 - calls wg.Done() when finished
	for _, num := range nums {
		sem <- struct{}{}
		wg.Add(1)

		go func(value int) {
			defer wg.Done()
			defer func() { <-sem }()

			results <- taskValue(value)
		}(num)
	}

	// Wait for all tasks to finish (wg.Wait())
	// close the results channel (close(results))
	// wait for the summing goroutine to finish (sumWg.Wait())
	// return the sum
	wg.Wait()
	close(results)
	sumWg.Wait()
	return sum
}
