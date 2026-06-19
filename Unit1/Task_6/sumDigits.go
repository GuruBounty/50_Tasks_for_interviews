package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(in, &num)
	fmt.Println(sumDigits(num))

}

func sumDigits(num int) int {
	if num < 0 {
		num = -num
	}
	result := 0
	for num > 0 {
		result += num % 10
		num /= 10
	}
	return result
}
