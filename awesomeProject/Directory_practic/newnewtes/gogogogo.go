package main

import "fmt"

func main() {
	var mapa map[int]int
	mapa[0] = 1
	mapa[5] = 0
	mapa[2] = 3
	for i, value := range mapa {
		fmt.Println(i, value)
	}
}
