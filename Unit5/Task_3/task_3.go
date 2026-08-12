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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	fmt.Println(ThreeSumEqualZero(a))
}

// Figure out there are three distinct elements whose sum is 0.
func ThreeSumEqualZero(a []int) string {
	if len(a) == 0 || len(a) < 3 {
		return "NO"
	}
	sort.Ints(a)

	for i := 0; i < len(a)-2; i++ {
		if a[i] > 0 {
			break
		}

		if i > 0 && a[i] == a[i-1] {
			continue
		}
		left := i + 1
		right := len(a) - 1

		for left < right {
			sum := a[i] + a[left] + a[right]
			if sum == 0 {
				return "YES"
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return "NO"
}
