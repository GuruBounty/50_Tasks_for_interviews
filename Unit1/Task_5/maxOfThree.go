package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var num1, num2, num3 int
	fmt.Fscan(in, &num1, &num2, &num3)

	fmt.Println(maxOfThree(num1, num2, num3))

}

func maxOfThree(num1, num2, num3 int) int {
	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}

	return max
}
