package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n, k int

	if _, err := fmt.Fscan(in, &m, &n, &k); err != nil {
		return
	}
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &nums[i])
	}
	cnt, sum := slove(nums, n, k)
	fmt.Printf("%d %d", cnt, sum)
}

func slove(nums []int, nWorkers, kCancel int) (int64, int64) {
	// - nums: значения тиков (длина M)
	// - nWorkers: количество воркеров
	// - kCancel: после отправки K тиков нужно вызвать cancel и остановить поток
	//
	// Нужно вернуть:
	// - сколько тиков реально обработали
	// - сумму квадратов обработанных тиков

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks := make(chan int, 64)
	results := make(chan int64, 64)

	var wg sync.WaitGroup
	wg.Add(nWorkers)

	for i := 0; i < nWorkers; i++ {
		go func() {
			defer wg.Done()
			for x := range tasks {
				results <- taskValue(x)
			}
		}()
	}

	go func() {
		defer close(tasks)

		limit := kCancel
		if limit > len(nums) {
			limit = len(nums)
		}
		if limit < 0 {
			limit = 0
		}

		for i := 0; i < limit; i++ {
			select {
			case <-ctx.Done():
				return
			case tasks <- nums[i]:
			}
		}
		cancel()

	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var cnt int64
	var sum int64
	for v := range results {
		cnt++
		sum += v
	}
	return cnt, sum
}

func taskValue(x int) int64 {
	return int64(x) * int64(x)
}
