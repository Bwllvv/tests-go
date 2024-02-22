package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {
	slice := []person{1, 2, 3, 4, 5}
	lev := person{
		id:   1,
		name: "lev",
	}
	fmt.Printf("Id: %d \nName: %s", lev.id, lev.name)
}
