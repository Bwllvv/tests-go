package main

import "fmt"

// пример стурктуры
type Pointsss struct {
	x int
	y int
}

// метод стуктуры
func (c Pointsss) methd() {
	fmt.Println(c.x, "икс вызов метода структуры")
}

// мэйн
func main() {
	fmt.Println("ad")
	pointers()
	structs()
	p2 := Pointsss{
		x: 2,
		y: 3,
	}
	p2.methd()
}

// функция с присваиванием значений стуктуры
func structs() {
	p1 := Pointsss{
		x: 1,
		y: 2,
	}
	fmt.Println(p1.x, p1.y, "точки заданные в стурктуре p1")
}

// ссылки на значения (экономия памяти)
func pointers() {
	a := 42
	b := 2
	p := &a
	fmt.Println(a, "значение а")
	fmt.Println(b, "значение в")
	fmt.Println(*p, "значение p (значение а через ссылку)", p, "ссылка на значение")

}
