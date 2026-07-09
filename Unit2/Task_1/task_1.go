package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()

	scanner.Scan()
	line2 := scanner.Text()

	if line1 == "" {
		fmt.Println(-1)
		return
	}
	strNumber := strings.Fields(line1)
	arr := make([]int, len(strNumber))
	for i, v := range strNumber {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(-1)
			return
		}
		arr[i] = num
	}
	target, err := strconv.Atoi(line2)
	if err != nil {
		fmt.Println(-1)
		return
	}
	index := parallelSearch(arr, target)
	fmt.Println(index)

}

func parallelSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	bufCh := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go searchInChunk(left, target, 0, &wg, bufCh)
	go searchInChunk(right, target, mid, &wg, bufCh)

	wg.Wait()
	close(bufCh)

	result := -1
	for idx := range bufCh {
		if idx >= 0 {
			if result == -1 || idx < result {
				result = idx
			}
		}
	}

	return result
}

func searchInChunk(chunk []int, target, startIndex int, wg *sync.WaitGroup, resChan chan int) {
	defer wg.Done()

	for i, v := range chunk {
		if v == target {
			resChan <- startIndex + i
			return
		}
	}
	resChan <- -1

}
