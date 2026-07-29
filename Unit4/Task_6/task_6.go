package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	result := mergeSorted(a, b)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, x := range result {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)
}

// merges two non-descending sorted arrays A and B into a single sorted array
func mergeSorted(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	for j < len(b) {
		res = append(res, b[j])
		j++
	}

	return res
}
