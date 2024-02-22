package main

import (
	"fmt"
	"sort"
)

func main() { // slice динамический массив + - элементы
	slice := []int{1, 23125, 7, 8}

	slice = append(slice, 0)

	slice[0] = 11

	sort.Ints(slice)
	fmt.Println(slice)

	// сортировка стринг слайса
	slices := []string{"abob", "bob", "coc", "doc"}
	sort.Strings(slices)
	fmt.Println(slices)

	//удаление эмемента из слайса
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// 1. Копировать последний элемент в индекс i.
	a[i] = a[len(a)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// 3. Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a) // [A B E D]
}
