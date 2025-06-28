package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("data not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Nafi" {
		return NotFoundError
	}

	// success

	return nil
}

func main() {
	err := GetById("")
	// err := GetById("Budi")
	// err := GetById("Nafi")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
