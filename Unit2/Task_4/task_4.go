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
	input := scanner.Text()

	strNumbers := strings.Fields(input)
	nums := make([]int, len(strNumbers))
	for i, v := range strNumbers {
		num, _ := strconv.Atoi(v)
		nums[i] = num
	}

	if len(nums) == 0 {
		fmt.Println(0)
		return
	}

	chunkSize := len(nums) / 3
	chunkA := nums[:chunkSize]
	chunkB := nums[chunkSize : 2*chunkSize]
	chunkC := nums[2*chunkSize:]

	mapChan := make(chan int, 3)
	reduceChan := make(chan int)

	var wgMap sync.WaitGroup
	wgMap.Add(3)

	var wgRed sync.WaitGroup
	wgRed.Add(1)

	go reduce(mapChan, reduceChan, &wgRed)

	go mapper(chunkA, mapChan, &wgMap)
	go mapper(chunkB, mapChan, &wgMap)
	go mapper(chunkC, mapChan, &wgMap)

	wgMap.Wait()
	close(mapChan)

	total := <-reduceChan
	wgRed.Wait()
	fmt.Println(total)
}

func mapper(chunk []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result int
	for _, v := range chunk {
		n := v * v
		result += n
	}
	out <- result
}

func reduce(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	total := 0
	for sum := range in {
		total += sum
	}
	out <- total
}
