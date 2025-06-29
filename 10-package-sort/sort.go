package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// cara standar untuk swap
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	// cara golang
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Nafi", 30},
		{"Furqon", 35},
		{"Diani", 25},
		{"Muhammad", 20},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
