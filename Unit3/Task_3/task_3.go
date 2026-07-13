package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	fmt.Print(longestUniqueSubstringLen(line))
}

func longestUniqueSubstringLen(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, maxLen := 0, 0
	runeS := []rune(s)
	symbols := make(map[rune]int)

	for right, value := range runeS {
		if lastIdx, ok := symbols[value]; ok && lastIdx >= left {
			left = lastIdx + 1
		}

		symbols[value] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
