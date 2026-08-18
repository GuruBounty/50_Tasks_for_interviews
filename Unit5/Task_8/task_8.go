package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	stack(in, out)
	_ = out.Flush()
}

func stack(in *bufio.Reader, out *bufio.Writer) {
	inStack := make([]int, 0)
	outStack := make([]int, 0)
	transfer := func() {
		for len(inStack) > 0 {
			outStack = append(outStack, inStack[len(inStack)-1])
			inStack = inStack[:len(inStack)-1]
		}
	}

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "push":
			var x int
			fmt.Fscan(in, &x)
			inStack = append(inStack, x)
			_ = x
		case "pop":
			if len(outStack) == 0 {
				transfer()
			}
			if len(outStack) == 0 {
				fmt.Fprintln(out, "ERROR")
			} else {
				val := outStack[len(outStack)-1]
				outStack = outStack[:len(outStack)-1]
				fmt.Fprintln(out, val)
			}
		case "front":
			if len(outStack) == 0 {
				transfer()
			}
			if len(outStack) == 0 {
				fmt.Fprintln(out, "ERROR")
			} else {
				val := outStack[len(outStack)-1]
				fmt.Fprintln(out, val)
			}
		case "size":
			size := len(inStack) + len(outStack)
			fmt.Fprintln(out, size)
		case "empty":
			if len(inStack)+len(outStack) == 0 {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		default:
			fmt.Fprintln(out, "ERROR")
		}
	}

}
