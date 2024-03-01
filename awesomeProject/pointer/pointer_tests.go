package main

import "fmt"

func main() {
	a := 0

	other(&a)
	fmt.Println(a)
}

func other(a *int) {
	b := 4
	fmt.Println(a)
	*a = b
}
