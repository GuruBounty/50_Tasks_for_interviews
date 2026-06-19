package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrDividByZero = errors.New("Divided by zero")

func saveDivid(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDividByZero
	}
	return a / b, nil
}
func main() {
	in := bufio.NewReader(os.Stdin)
	var num1, num2 int
	fmt.Fscan(in, &num1, &num2)
	result, err := saveDivid(num1, num2)
	if err != nil {
		fmt.Println("ERR")
		return
	}
	fmt.Println(result)
}
