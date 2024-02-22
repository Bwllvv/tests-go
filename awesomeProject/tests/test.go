package main

import "fmt"

func main() {

	var d1, d2, d3 int
	var do string

	fmt.Println("Первое число")
	fmt.Scanf("%d\n", &d1)
	fmt.Println("Второе число")
	fmt.Scanf("%d\n", &d2)
	fmt.Println("Действие")
	fmt.Scanf("%s\n", &do)

	switch do {
	case "plus":
		d3 = d1 + d2
		fmt.Println(d3)
	case "minus":
		d3 = d1 - d2
		fmt.Println(d3)
	case "multiply":
		d3 = d1 * d2
	case "del":
		d3 = d1 / d2
	}

}
