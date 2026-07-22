package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()

	scanner.Scan()
	line2 := scanner.Text()

	scanner.Scan()
	line3 := scanner.Text()

	if line1 == "" || line2 == "" || line3 == "" {
		fmt.Println("")
		return
	}
	CntElemts, err := strconv.Atoi(line1)
	if err != nil || CntElemts <= 0 {
		fmt.Println("")
		return
	}
	idx, err := strconv.Atoi(line3)
	if err != nil || idx < 0 || idx >= CntElemts {
		fmt.Printf("")
		return
	}
	elemts := strings.Split(line2, " ")

	fmt.Print(RemoveElemtInSlice(elemts, idx))
}

func RemoveElemtInSlice(elemts []string, idx int) string {
	// TODO: write a function that removes the element at index k and returns a new slice of length n-1.
	var newSlice []string

	for i := 0; i < len(elemts); i++ {
		if i != idx {
			newSlice = append(newSlice, elemts[i])
		}
	}
	result := strings.Join(newSlice, " ")
	return result
}
