package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var k int
	fmt.Fscan(in, &k)

	if n == 0 {
		fmt.Println(0)
		return
	}
	kthLagestElem := quickSelect(a, k)
	fmt.Println(kthLagestElem)
}

func quickSelect(a []int, k int) int {
	kSmall := len(a) - k
	rand.Seed(time.Now().UnixNano())
	l, r := 0, len(a)-1

	for l <= r {
		pivotIdx := l + rand.Intn(r-l+1)
		p := partition(a, l, r, pivotIdx)
		if p == kSmall {
			return a[p]
		} else if p < kSmall {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return 0

}

func partition(a []int, l, r, pivotIdx int) int {
	// TODO:
	// 1) Choose pivot = a[pivotIndex]
	pivot := a[pivotIdx]
	// 2) Change a[pivotIndex] to a[r], that  pivot temporarlly found itself at the end
	a[pivotIdx], a[r] = a[r], a[pivotIdx]
	// 3) Create bookmark store = l
	store := l
	// 4) Go i from l to r-1:
	for i := l; i < r; i++ {
		//    if a[i] < pivot:
		//       swap(a[i], a[store]); store++
		if a[i] < pivot {
			a[i], a[store] = a[store], a[i]
			store++
		}
	}
	// 5)  swap(a[store], a[r])
	a[store], a[r] = a[r], a[store]
	// 6) вернуть store
	return store
}
