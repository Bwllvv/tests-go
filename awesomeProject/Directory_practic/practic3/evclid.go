package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(check(arr))

}

func check(arr []int) int {
	var newarr []string
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for idx := 0; idx < len(arr)-1; idx++ {
		if arr[idx] == arr[idx+1] {

		} else {
			newarr = append(newarr, strconv.Itoa(arr[idx]))
			//fmt.Println(newarr)
			count++
		}
	}
	if arr[len(arr)-2] != arr[len(arr)-1] {
		newarr = append(newarr, strconv.Itoa(arr[len(arr)-1]))
	}
	for len(newarr) != len(arr) {
		newarr = append(newarr, "_")
	}
	fmt.Println(newarr)
	return count

}
