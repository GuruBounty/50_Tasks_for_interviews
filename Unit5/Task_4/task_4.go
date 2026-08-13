package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	intervals := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		//var l, r int
		//fmt.Fscan(in, &l, &r)
		//intervals = append(intervals, [2]int{l, r})
		fmt.Fscan(in, &intervals[i][0], &intervals[i][1])
	}

	merged := mergeIntervals(intervals)

	fmt.Println(len(merged))
	for _, it := range merged {
		fmt.Println(it[0], it[1])
	}
}

func mergeIntervals(a [][2]int) [][2]int {
	if len(a) == 0 {
		return nil
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})

	res := make([][2]int, 0, len(a))
	cur := a[0]
	for i := 1; i < len(a); i++ {
		if a[i][0] <= cur[1] { // overlap or touching (closed intervals)
			if a[i][1] > cur[1] {
				cur[1] = a[i][1]
			}
		} else {
			res = append(res, cur)
			cur = a[i]
		}
	}
	res = append(res, cur)
	return res
}
