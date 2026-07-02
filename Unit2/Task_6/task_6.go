package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем M и N:
	// M — количество задач (чисел)
	// N — количество воркеров
	var m, n int
	if _, err := fmt.Fscan(in, &m, &n); err != nil {
		return
	}

	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &nums[i])
	}

	ans := solve(nums, n)
	fmt.Print(ans)
}

func solve(nums []int, nWorkers int) int64 {
	// solve — главная функция, которую вы должны реализовать.
	//
	// Нам дано:
	// - nums: список задач (числа)
	// - nWorkers: количество воркеров (горутины-работники)
	//
	// Нужно:
	// 1) Создать канал задач с буфером РОВНО 10.
	tasks := make(chan int, 10)
	// 2) Запустить nWorkers воркеров. Каждый воркер:
	//    - читает числа из канала tasks (range по каналу)
	//    - для каждого числа считает taskValue(x)
	//    - отправляет результат в канал results
	results := make(chan int64, len(nums))
	var wg sync.WaitGroup
	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			defer wg.Done()
			for task := range tasks {
				val := taskValue(task)
				results <- val
			}
		}()
	}
	// 3) Отправить все числа nums в канал tasks.
	for _, num := range nums {
		tasks <- num
	}
	// 4) Закрыть канал tasks (чтобы воркеры поняли, что задач больше нет).
	close(tasks)
	// 5) Корректно завершить воркеров (WaitGroup).
	// 6) Корректно закрыть results (только ПОСЛЕ того, как воркеры закончили писать).
	go func() {
		wg.Wait()
		close(results)
	}()
	// 7) Собрать все результаты из results и вернуть их сумму.
	//
	var sum int64 = 0
	// Нельзя:
	// - закрывать results раньше времени (будет panic)
	// - оставлять воркеров висеть (дедлок)
	// - терять задачи
	for v := range results {
		sum += v
	}
	return sum
}

func taskValue(task int) int64 {
	return int64(task) * int64(task)
}
