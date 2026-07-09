package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nStr := scanner.Text()

	scanner.Scan()
	timeStr := scanner.Text()

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		fmt.Println("Неверный ввод")
		return
	}

	workTimeMs, err := strconv.Atoi(timeStr)
	if err != nil || workTimeMs <= 0 {
		fmt.Println("Неверный ввод")
		return
	}
	workTime := time.Duration(workTimeMs) * time.Microsecond

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, workTime, &wg)
	}

	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}

func worker(id int, workTime time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала работу\n", id)
	time.Sleep(workTime)
	fmt.Printf("Горутина  %d завершила работу\n", id)
}
