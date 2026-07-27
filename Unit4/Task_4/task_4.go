package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &words[i])
	}
	groups := groupAnagrams(words)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, g := range groups {
		for i, w := range g {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, w)
		}
		fmt.Fprintln(out)
	}
}

func groupAnagrams(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}

	groups := make([][]string, 0, len(words))
	idx := make(map[string]int, len(words))

	for _, word := range words {
		key := canonicalKey(word)
		if index, ok := idx[key]; ok {
			groups[index] = append(groups[index], word)
		} else {
			idx[key] = len(groups)
			groups = append(groups, []string{word})
		}
	}
	return groups
}

func canonicalKey(s string) string {
	s = strings.ToLower(s)
	counts := make([]int, 26)
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			counts[ch-'a']++
		}
	}

	var b []byte
	for i, count := range counts {
		for j := 0; j < count; j++ {
			b = append(b, byte('a'+i))
		}
	}
	return string(b)
}
