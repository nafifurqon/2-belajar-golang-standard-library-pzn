package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`strings.Contains("Nafi Furqon", "Nafi")`, strings.Contains("Nafi Furqon", "Nafi"))
	fmt.Println(`strings.Split("Nafi Furqon", " ")`, strings.Split("Nafi Furqon", " "))
	fmt.Println(`strings.ToLower("Nafi Furqon")`, strings.ToLower("Nafi Furqon"))
	fmt.Println(`strings.ToUpper("Nafi Furqon")`, strings.ToUpper("Nafi Furqon"))
	fmt.Println(`strings.Trim("     Nafi Furqon     ", " ")`, strings.Trim("     Nafi Furqon     ", " "))
	fmt.Println(`strings.ReplaceAll("Nafi Furqon Nafi Diani", "Nafi", "Budi")`, strings.ReplaceAll("Nafi Furqon Nafi Diani", "Nafi", "Budi"))
	fmt.Println(`strings.Join([]string{"Muhammad", "Nafi", "Furqon"}, " ")`, strings.Join([]string{"Muhammad", "Nafi", "Furqon"}, " "))
}
