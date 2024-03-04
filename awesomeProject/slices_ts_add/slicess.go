package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 5, 10)
	fmt.Println(arr, len(arr), cap(arr))
	change_id(arr)
	fmt.Println(arr)
	appendd(arr)
	fmt.Println(1111, arr, len(arr), cap(arr))
}
func appendd(arr []int) {
	arr = append(arr, 50)
	//arr[0] = 2000
	fmt.Println(arr, len(arr), cap(arr))
}

func change_id(arr []int) {
	arr[2] = 55
}
