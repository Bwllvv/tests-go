package main

import "fmt"

func remove[T comparable](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1] // Заменяем удаляемый элемент на последний
	return slice[:len(slice)-1]    // Усекаем слайс, исключая последний элемент
}

func main() {
	beforeRemoveSlice := []int{1, 2, 3, 4, 5}
	indexRemove := 2

	afterRemoveSlice := remove(beforeRemoveSlice, indexRemove)
	fmt.Println("after remove int slice:", afterRemoveSlice)
}
