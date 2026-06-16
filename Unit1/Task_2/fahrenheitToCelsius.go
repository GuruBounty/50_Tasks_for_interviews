package main

import (
	"fmt"
	"strconv"
)

func FahrenheitToCelsius(f any) (c float64, err error) {
	switch v := f.(type) {
	case float64:
		c = (v - 32) * 5 / 9
		return c, nil
	default:
		return 0, fmt.Errorf("invalid type")
	}

}

func main() {
	var input string
	fmt.Scanln(&input)

	if f, err := strconv.ParseFloat(input, 64); err == nil {
		c, err := FahrenheitToCelsius(f)
		if err != nil {
			fmt.Println("invalid input")
		} else {
			fmt.Println("%.2", c)
		}
	} else {
		_, err := FahrenheitToCelsius(input)
		if err != nil {
			fmt.Println("invalid input")
		}
	}
}
