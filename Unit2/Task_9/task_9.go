package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var B, S, v int
	if _, err := fmt.Fscan(in, &B, &S, &v); err != nil {
		return
	}
	// A buffered channel of a specified capacity
	ch := make(chan int, B)

	// pre-fill channel S with numbers
	for i := 0; i < S; i++ {
		var x int
		fmt.Fscan(in, &x)
		ch <- x
	}

	// Try to send v
	sendOk := 0
	if TrySend(ch, v) {
		sendOk = 1
	}

	// Try to read one value
	recvVal := -1
	if x, ok := TryRecv(ch); ok {
		recvVal = x
	}
	fmt.Printf("%d %d", sendOk, recvVal)
}

func TryRecv(ch chan int) (int, bool) {
	select {
	case x := <-ch:
		return x, true
	default:
		return 0, false
	}
}

func TrySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}

}
