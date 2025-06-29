package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println("slices.Min(names)", slices.Min(names))
	fmt.Println("slices.Min(values)", slices.Min(values))
	fmt.Println("slices.Max(names)", slices.Max(names))
	fmt.Println("slices.Max(values)", slices.Max(values))
	fmt.Println(`slices.Contains(names, "Nafi")`, slices.Contains(names, "Nafi"))
	fmt.Println(`slices.Index(names, "Nafi")`, slices.Index(names, "Nafi"))
	fmt.Println(`slices.Index(names, "Paul")`, slices.Index(names, "Paul"))
}
