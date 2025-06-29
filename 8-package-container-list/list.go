package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Muhammad")
	data.PushBack("Nafi")
	data.PushBack("Furqon")
	data.PushBack("Diani")

	var head *list.Element = data.Front()
	fmt.Println("head.Value", head.Value) // Muhammad

	next := head.Next()
	fmt.Println("next.Value", next.Value) // Nafi

	next = next.Next()
	fmt.Println("next.Value", next.Value) // Furqon

	next = next.Next()
	fmt.Println("next.Value", next.Value) // Furqon

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("e.Value", e.Value)
	}

}
