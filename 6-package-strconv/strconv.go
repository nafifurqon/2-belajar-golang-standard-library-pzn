package main

import (
	"fmt"
	"strconv"
)

func main() {
	// result, err := strconv.ParseBool("SALAH") // Error strconv.ParseBool: parsing "SALAH": invalid syntax
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", result)
	}

	// resultInt, err := strconv.Atoi("SALAH") // Error strconv.Atoi: parsing "SALAH": invalid syntax
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println("Binary", binary)

	stringInt := strconv.Itoa(999)
	fmt.Println("String", stringInt)
}
