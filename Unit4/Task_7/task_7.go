package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	b := make([]int, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	res := intersect(a, b)

	out := bufio.NewWriter(os.Stdin)
	defer out.Flush()

	for i, x := range res {
		if i > 0 {
			fmt.Fprintf(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)
}

func intersect(a, b []int) []int {
	//TODO:
	// 1. Create setA := map[int]struct{}{} and put all elements a
	setA := make(map[int]struct{})
	for _, x := range a {
		setA[x] = struct{}{}
	}
	// 2. Create setRes for result
	setRes := make(map[int]struct{})
	// 3. Iterate over b and check if element exists in setA, if yes, add to setRes
	for _, x := range b {
		if _, ok := setA[x]; ok {
			setRes[x] = struct{}{}
		}
	}
	// 4. Move the setRes keys to the res section
	res := make([]int, 0, len(setRes))
	for x := range setRes {
		res = append(res, x)
	}

	// 5. Sort the res slice
	//slices.Sort(res)
	//sort.Ints(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]

	})

	return res
}
