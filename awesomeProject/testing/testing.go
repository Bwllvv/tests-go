package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := 32
	a = sum(a)
	println("hello, my name is:", a, "\nare you?")
	slice := []int{12, 2, 3, 4, 5, 6}
	slice[2] = 0
	for i := 0; i < len(slice); i++ {
		for j := i; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[j]
				slice[j] = slice[i]
				slice[i] = temp
			}
		}
	}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	var x int

	var y int
	x = 3
	y = 3
	slice_double := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			print(slice_double[i][j])
			print(" ||| ")
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	// Читаем строку
	if scanner.Scan() {
		// Получаем считанную строку
		input := scanner.Text()
		if isNumeric(input) {
			fmt.Println("Вы ввели:", input)
		} else {
			println("n/a")
		}
	} else {
		fmt.Println("Ошибка чтения ввода")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		println("okay")
	}
}

func sum(a int) int {
	a += a
	return a
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
