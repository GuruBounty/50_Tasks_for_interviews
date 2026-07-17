package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	input := string(data)

	fmt.Print(normalizeSpaces(input))
}

// normalizeSpaces collapses any sequence of Unicode spaces into a single regular space
func normalizeSpaces(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	prevSpace := false
	for _, v := range s {
		if unicode.IsSpace(v) {
			if !prevSpace {
				b.WriteByte(' ')
				prevSpace = true
			}
		} else {
			b.WriteRune(v)
			prevSpace = false
		}
	}
	return b.String()
}
