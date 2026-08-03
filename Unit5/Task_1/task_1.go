package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var x int
	fmt.Fscan(in, &x)

	fmt.Println(findIndex(a, x))
}

func findIndex(a []int, x int) int {
	if len(a) == 0 {
		return -1
	}
	left := 0
	rigth := len(a) - 1

	for left <= rigth {
		mid := (left + rigth) / 2

		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			left = mid + 1
		} else {
			rigth = mid - 1
		}
	}
	return -1
}
