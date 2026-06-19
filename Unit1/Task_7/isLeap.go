package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var year int
	fmt.Fscan(in, &year)

	if isLeap(year) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isLeap(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
