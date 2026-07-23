package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	res := dedupPreserveOrder(a)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, s := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, s)
	}
	fmt.Fprintln(out)
}

func dedupPreserveOrder(a []string) []string {
	// TODO: create a set for the strings you've already encounterd
	// Hint: make(map[string]struct{}, len(a))
	// TODO: set the record index w (write index), to 0 initially.
	// TODO: iterate through all elements of the input slice form left ot right:
	// - if the element is already in the set, skip it
	// - otherwise:
	// * add it to the set
	// * store it in a[w]
	// * increment w
	// TODO: (recommended): clean the tail a[w:] (assign ""), so as not to hold onto links
	// TODO: return a[:w]
	set := make(map[string]struct{}, len(a))
	w := 0
	for _, s := range a {
		if _, ok := set[s]; !ok {
			set[s] = struct{}{}
			a[w] = s
			w++
		}
	}
	//a[w:] = ""
	return a[:w]
}
