package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n')
	p, _ := in.ReadString('\n')

	s = strings.TrimRight(s, "\r\n")
	p = strings.TrimRight(p, "\r\n")

	fmt.Print(prefixSuffixStatus(s, p))
}

func prefixSuffixStatus(s, p string) string {
	// TODO : two string (UTF-8) need to check if p is a prefix and/or suffix of the sting s
	// wihout strings.HasPrefix/HasSuffix
	// Unicode-safe: comparosion by runes
	// PREFIX if p is prefix s
	// SUFFIX if p is suffix s
	// BOTH if both
	// NONE if none
	runeS, runeP := []rune(s), []rune(p)
	if len(runeP) > len(runeS) {
		return "NONE"
	}
	if len(runeP) == 0 {
		return "BOTH"
	}

	isPrefix := true
	for i := 0; i < len(runeP); i++ {
		if runeS[i] != runeP[i] {
			isPrefix = false
			break
		}
	}

	isSuffix := true
	for i := 0; i < len(runeP); i++ {
		if runeP[i] != runeS[len(runeS)-len(runeP)+i] {
			isSuffix = false
			break
		}
	}

	switch {
	case isPrefix && isSuffix:
		return "BOTH"
	case isPrefix:
		return "PREFIX"
	case isSuffix:
		return "SUFFIX"
	default:
		return "NONE"
	}
}
