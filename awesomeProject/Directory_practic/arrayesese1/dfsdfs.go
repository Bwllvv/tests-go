package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
