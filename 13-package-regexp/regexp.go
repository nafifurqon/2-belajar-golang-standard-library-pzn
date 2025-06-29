package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`n[a-z]{2}i`)

	fmt.Println(`regex.MatchString("nafi")`, regex.MatchString("nafi"))
	fmt.Println(`regex.MatchString("nabi")`, regex.MatchString("nabi"))
	fmt.Println(`regex.MatchString("nARi")`, regex.MatchString("nARi"))

	fmt.Println(regex.FindAllString("nafi nabi nARi niki n8n", 10))
}
