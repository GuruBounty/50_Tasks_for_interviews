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

	fields := parseCSVLine(line)
	for _, f := range fields {
		fmt.Println(f)
	}
}

func parseCSVLine(line string) []string {
	// simplified CSV parser
	// Delimiter: ','
	// Quotation marks: '"', no escaping
	// Remove the outer quotation marks but keep the content inside
	if len(line) == 0 {
		return nil
	}
	lineRune := []rune(line)
	field := ""
	fields := []string{}
	for _, v := range lineRune {
		if v == '"' {
			continue
		}
		if v == ',' {
			fields = append(fields, field)
			field = ""
			continue
		}
		field += string(v)
	}
	fields = append(fields, field)
	return fields
}
