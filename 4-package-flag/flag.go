package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)

	/*
		input = go run 4-package-flag/flag.go -username=Nafi -password="Rahasia 123" -host=192.168.0.1 -port=3000
		output =
			username Nafi
			password Rahasia 123
			host 192.168.0.1
			port 3000
	*/
}
