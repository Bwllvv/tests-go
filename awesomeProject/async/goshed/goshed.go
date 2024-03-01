package main

import "fmt"

func main() {
	arr := []int{1, 22, 2, 2, 1, 3, 1, 1}
	ch := make(chan int)
	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

}

func sum(arr []int, ch chan int) {
	suma := 0
	for i := 0; i < len(arr); i++ {
		suma += arr[i]
	}
	ch <- suma
}
