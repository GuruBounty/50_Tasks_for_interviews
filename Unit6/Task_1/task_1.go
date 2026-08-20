package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrInvalidID = errors.New("invalid id")
var ErrStorage = errors.New("storage error")

func loadUser(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}
	if id == 13 {

		return fmt.Errorf("%w", ErrStorage)
	}
	return nil
}
func main() {
	in := bufio.NewReader(os.Stdin)

	var id int
	fmt.Fscan(in, &id)

	err := loadUser(id)

	a := 0
	if errors.Is(err, ErrInvalidID) {
		a = 1
	}
	b := 0
	if errors.Is(err, ErrStorage) {
		b = 1
	}
	c := 0
	if err == ErrStorage {
		c = 1
	}

	fmt.Printf("%d %d %d", a, b, c)
}
