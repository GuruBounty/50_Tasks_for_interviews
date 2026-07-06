package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	// nA, nB, nC - how many numbers in every channel
	var nA, nB, nC int

	if _, err := fmt.Fscan(in, &nA, &nB, &nC); err != nil {
		return
	}
	chanA := make(chan int)
	chanB := make(chan int)
	chanC := make(chan int)

	// fanIn runs its gorutines and will output to out
	out := fanIn(chanA, chanB, chanC)

	var sum int64
	var sumWg sync.WaitGroup
	sumWg.Add(1)
	//send number in a/b/c.
	go func() {
		defer sumWg.Done()
		for v := range out {
			sum += int64(v)
		}
	}()

	//filling channel A
	for i := 0; i < nA; i++ {
		var x int
		fmt.Fscan(in, &x)
		chanA <- x
	}
	//close channel A this's "end of data" signal
	close(chanA)

	//filling channel B
	for i := 0; i < nB; i++ {
		var x int
		fmt.Fscan(in, &x)
		chanB <- x
	}
	close(chanB)

	//filling channel C
	for i := 0; i < nC; i++ {
		var x int
		fmt.Fscan(in, &x)
		chanC <- x
	}
	close(chanC)

	//Wait,until:
	// - fanIn has finished reading all channels
	// -out has been closed
	// -the summation has finished
	sumWg.Wait()

	fmt.Print(sum)
}

// fanIn - merge three input channels into a single output channel
func fanIn(chanA, chanB, chanC <-chan int) <-chan int {
	//We are given three channels chanA, chanB, chanC (each of which will eventually be closed)
	// We need to return a channel out, that outputs ALL values form a,b,c (in any order)

	//Key idea:
	//-For each input channel we start a forwarder goroutine: it reads the input and writes to out
	//-The WaitGroup waits for all forwarders to finish.
	//-After wg.Wait(), we close out (it must be closed exactly once, in one place).

	out := make(chan int)

	//WaitGroup counts how many goroutines are still working
	var wg sync.WaitGroup
	wg.Add(3)

	//forward reads all values from a certain channel and sends them to out
	forward := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			out <- v
		}
	} // the range will end when ch is closed

	//run the loop on each source channel
	go forward(chanA)
	go forward(chanB)
	go forward(chanC)

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
