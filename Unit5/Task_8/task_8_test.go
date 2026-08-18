package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runStack(input string) string {
	in := bufio.NewReader(strings.NewReader(input))
	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	stack(in, out)
	_ = out.Flush()
	return buf.String()
}

func TestStackUsingTwoQueues(t *testing.T) {
	input := "9\npush 1\npush 2\nfront\npop\nfront\npop\npop\nempty\nsize\n"
	want := "1\n1\n2\n2\nERROR\nYES\n0\n"

	got := runStack(input)
	if got != want {
		t.Fatalf("unexpected output\nwant:\n%s\ngot:\n%s", want, got)
	}
}
