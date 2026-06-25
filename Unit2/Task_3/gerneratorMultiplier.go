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
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		nums[i] = num
	}
	outGen := make(chan int, len(nums))
	outMult := make(chan int, len(nums))

	var wg sync.WaitGroup
	wg.Add(3)

	go generator(nums, outGen, &wg)
	go multiplier(outGen, outMult, &wg)
	go output(outMult, &wg)

	wg.Wait()
}

func output(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range in {
		fmt.Printf("Результат: %d\n", num+10)
	}
}

func multiplier(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for num := range in {
		out <- num * 2
	}
}

func generator(nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for _, v := range nums {
		out <- v
	}
}
