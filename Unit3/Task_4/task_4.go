package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	fmt.Print(rleCompressRunes(line))
}

func rleCompressRunes(s string) string {
	// TO DO: RLE - compress
	// Format every group = rune + count; aaabb → a3b2
	if len(s) == 0 {
		return ""
	}

	runeS := []rune(s)
	var builder strings.Builder

	count := 1
	for i := 1; i < len(runeS); i++ {
		if runeS[i] == runeS[i-1] {
			count++
			continue
		}

		builder.WriteRune(runeS[i-1])
		builder.WriteString(strconv.Itoa(count))
		count = 1
	}

	builder.WriteRune(runeS[len(runeS)-1])
	builder.WriteString(strconv.Itoa(count))

	return builder.String()
}
