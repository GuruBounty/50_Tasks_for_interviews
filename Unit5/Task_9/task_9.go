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

	var K int
	fmt.Fscan(in, &K)
	fmt.Println(sumSubSlice(a, K))
}

func sumSubSlice(a []int, K int) string {
	seen := make(map[int]struct{}, len(a)+1)
	seen[0] = struct{}{}

	pref := 0

	for i := 0; i < len(a); i++ {
		pref += a[i]
		if _, ok := seen[pref-K]; ok {
			return "YES"
		}
		seen[pref] = struct{}{}
	}
	return "NO"
}
