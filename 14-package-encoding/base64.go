package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Muhammad Nafi Furqon Diani"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("encoded", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("decoded", string(decoded)) // ubah jadi string karena return aslinya berupa byte
	}
}
