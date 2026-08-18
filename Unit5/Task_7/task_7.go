package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	line, _ := in.ReadString('\n')
	if len(line) > 0 && line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	if len(line) > 0 && line[len(line)-1] == '\r' {
		line = line[:len(line)-1]
	}
	fmt.Println(braces(line))
}

func braces(line string) string {
	stack := make([]byte, 0, len(line))
	matches := func(open, close byte) bool {
		if open == '(' && close == ')' {
			return true
		}
		if open == '[' && close == ']' {
			return true
		}
		if open == '{' && close == '}' {
			return true
		}
		return false
	}

	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return "NO"
			} else if !matches(stack[len(stack)-1], ch) {
				return "NO"
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		_ = ch
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}
