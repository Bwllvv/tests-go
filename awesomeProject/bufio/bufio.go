package main

import "fmt"

type person struct {
	id    int
	name  string
	score []int
	class
}

type class struct {
	clan  string
	rassa string
}

func (p *person) addScore(a int) {
	p.score = append(p.score, a)
}

func (p person) TakeAllOutput() {
	println(p.id)
	println(p.name)
	fmt.Println(p.score)
	fmt.Println("clan:", p.class.clan, "\nrassa:", p.class.rassa)
}

func main() {
	levichka := person{
		id:    1,
		name:  "leva",
		score: []int{1, 2, 33, 22, 42},
		class: class{clan: "daun",
			rassa: "pidoras",
		},
	}
	levichka.addScore(55)
	levichka.TakeAllOutput()
}
