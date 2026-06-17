package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(countVowels(input))
}

func countVowels(s string) int {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'а': true, 'е': true, 'ё': true, 'и': true, 'о': true,
		'у': true, 'ы': true, 'э': true, 'ю': true, 'я': true,
	}
	count := 0
	//str := unicode.ToLower(s)
	for _, v := range s {
		runeV := unicode.ToLower(v)
		if _, ok := vowels[runeV]; ok {
			count++
		}
	}
	return count
}
