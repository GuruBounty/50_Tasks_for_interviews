package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	s := string(data)

	val, ok := atoi32(s)
	if !ok {
		fmt.Print("ERROR")
		return
	}
	fmt.Print(val)
}

func atoi32(s string) (int32, bool) {
	// TODD Parse s into int32(base10) checking for overflow.
	// Don't use strconv.Atoi / strconv.ParseInt and other parsers
	// Return (value, true) if success else (0, false)
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}

	negative := false
	if s[0] == '+' || s[0] == '-' {
		if len(s) == 1 {
			return 0, false
		}
		if s[0] == '-' {
			negative = true
		}
		s = s[1:]
	}

	var value int32
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		digit := int32(c - '0')

		if negative {
			if value < (int32(math.MinInt32)+digit)/10 {
				return 0, false
			}
			value = value*10 - digit
		} else {
			if value > (math.MaxInt32-digit)/10 {
				return 0, false
			}
			value = value*10 + digit
		}
	}

	return value, true
}
