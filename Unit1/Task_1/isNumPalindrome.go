package main

import (
	"fmt"
	"strconv"
)

func isNumPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	temp := num
	reverse := 0

	for temp > 0 { //num = 121
		digit := temp % 10           //1.digit=1   2. digit=2 3. digit=1
		reverse = reverse*10 + digit //1. r=0*10+1 2. r=1*10+2=12 3. r=12*10+1=121
		temp /= 10                   //1. temp =12 2. temp = 1 3. temp =0

	}
	if reverse != num {
		return false
	}
	return true
}

func main() {
	var input string
	fmt.Scanln(&input)
	num, _ := strconv.Atoi(input)
	fmt.Println(isNumPalindrome(num))
}
