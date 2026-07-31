package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counter struct {
	Value int
}

func (c *Counter) Inc(delta int) {
	c.Value += delta
}

func (c Counter) Get() int {
	return c.Value
}
func main() {
	in := bufio.NewReader(os.Stdin)

	var q int
	if _, err := fmt.Fscan(in, &q); err != nil {
		return
	}

	var c Counter

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		if cmd == "INC" {
			var x int
			fmt.Fscan(in, &x)
			c.Inc(x)
		} else if cmd == "GET" {
			fmt.Fprintln(out, c.Get())
		}
	}
}
