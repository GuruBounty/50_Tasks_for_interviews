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

	if n == 0 {
		fmt.Println(0)
		return
	}
	minElment := findMinElement(a)
	fmt.Println(minElment)
}

func findMinElement(a []int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := l + (r-l)/2

		if a[mid] > a[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return a[l]
}
