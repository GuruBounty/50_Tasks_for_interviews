package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	t, _ := in.ReadString('\n')

	s = strings.TrimRight(s, "\r\n")
	t = strings.TrimRight(t, "\r\n")

	if isAnagramUnicodeCaseInsensitive(s, t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isAnagramUnicodeCaseInsensitive(s, t string) bool {
	// TO DO check anagram for runse + lower case
	// Key: map[rune]int, unicode.ToLower
	if len(s) != len(t) {
		return false
	}
	runeS := []rune(s)
	runeT := []rune(t)
	cnt := make(map[rune]int)

	for _, v := range runeS {
		cnt[unicode.ToLower(v)]++
	}
	for _, v := range runeT {
		lower := unicode.ToLower(v)
		if cnt[lower] == 0 {
			return false
		}
		cnt[lower]--
	}

	return true
}
