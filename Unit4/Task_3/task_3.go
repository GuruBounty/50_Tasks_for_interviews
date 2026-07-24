package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	if _, err := fmt.Fscan(in, &n, &k); err != nil {
		return
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	res := topKFrequent(a, k)

	out := bufio.NewWriter(os.Stdin)
	defer out.Flush()

	for i, x := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)

}

func topKFrequent(a []int, k int) []int {
	// TODO:
	// 1. If k==0 or len(a)==0 return nil slice
	if k == 0 || len(a) == 0 {
		return []int{}
	}
	//2. Count frequents use map[int]int:
	//freq:=make(map[int] int, ../)
	//freq[x]++
	freq := make(map[int]int)
	for _, x := range a {
		freq[x]++
	}

	// 3. Put freq in slice pair (value, count)
	pairs := make([][]int, 0, len(freq))
	for val, cnt := range freq {
		pairs = append(pairs, []int{val, cnt})
	}

	// 4. Order pairs:
	// - count desc
	// -value abs (tie-break)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] != pairs[j][1] {
			return pairs[i][1] > pairs[j][1]
		}

		return pairs[i][0] < pairs[j][0]
	})

	// 5. Take firs min(k, len(pairs)) values and return them
	result := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i][0])
	}
	return result

}
