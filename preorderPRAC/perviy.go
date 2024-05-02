package main

import (
	"fmt"
)

type worker struct {
	id   int
	name string
}

func main() {
	fmt.Println("hello World")
	arr := make([]int, 3, 4)
	change_id(arr)
	fmt.Println(arr, len(arr), cap(arr))
	appendd(arr)
	arr = arr[:4]
	fmt.Println(arr, len(arr), cap(arr))
	appendd(arr)
	//arr = arr[:5]
	fmt.Println(arr, len(arr), cap(arr))
	lev := worker{
		id:   10,
		name: "lev",
	}
	fmt.Println(lev.id)
}

func appendd(arr []int) {
	arr = append(arr, 50)
	fmt.Println(arr, len(arr), cap(arr))
}

func change_id(arr []int) {
	arr[2] = 55
}
