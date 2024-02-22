package main

import "fmt"

type interf interface {
	Solve() int
}

type structure struct {
	N1, N2 int
}

type otherstruct struct {
	a, b int
}

func (s *structure) Solve() int { // указана ссылка на структуру , а сам по себе это метод структуры
	return s.N1 + s.N2
}
func (o *otherstruct) Solve() int {
	return o.a * o.b
}

func main() {
	var a, b interf
	//sh := structure{1, 2}
	//os := otherstruct{2, 3}

	a = &structure{1, 2}
	b = &otherstruct{2, 3}
	fmt.Println(a.Solve())
	fmt.Println(b.Solve())
}
