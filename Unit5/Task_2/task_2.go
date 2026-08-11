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

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var target int
	fmt.Fscan(in, &target)

	fmt.Println(FoundTwoElemtsIfSumEqualTarget(nums, target))
}

func FoundTwoElemtsIfSumEqualTarget(nums []int, target int) string {
	if len(nums) == 0 {
		return "NO"
	}
	seen := make(map[int]bool)
	for _, v := range nums {
		if seen[target-v] {
			return "YES"
		} else {
			seen[v] = true
		}
	}
	return "NO"
}
