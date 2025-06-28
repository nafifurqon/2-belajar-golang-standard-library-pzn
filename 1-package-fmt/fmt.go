package main

import "fmt"

func main() {
	firstName := "Muhammad"
	lastName := "Nafi"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)

	age := 27
	fmt.Printf("Umur saya %d tahun\n", age)
	fmt.Printf("Umur saya dalam binary adalah %b tahun\n", age)
}
