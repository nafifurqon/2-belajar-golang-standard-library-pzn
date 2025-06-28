package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i, arg := range args {
		fmt.Printf("index ke %d = %s\n", i, arg)
	}

	/*
		input = go run 3-package-os/os.go Muhammad Nafi Furqon
		output =
			index ke 0 = /Users/muhammadnafifurqondiani/Library/Caches/go-build/0a/0a7acce419aea18b0647567aed37559c029f8b02947e2b855342d24bf84ab3b9-d/os
			index ke 1 = Muhammad
			index ke 2 = Nafi
			index ke 3 = Furqon
	*/

	/*
		input = go run 3-package-os/os.go "Muhammad Nafi Furqon"
		output =
			index ke 0 = /var/folders/19/xxlwxwd14qg7c25l_16j4mx40000gn/T/go-build757820562/b001/exe/os
			index ke 1 = Muhammad Nafi Furqon
	*/

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("hostname = %s\n", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
