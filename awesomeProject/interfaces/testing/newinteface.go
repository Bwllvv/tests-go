package main

import (
	"fmt"
)

type cooker interface {
	cook()
	getName() string
}

type brother struct {
	Namefull  string
	Older     int
	Breakfast string
}
type mother struct {
	Namefull  string
	Older     int
	Breakfast string
}

func (m *mother) cook() {
	fmt.Printf("done")
}
func (m *mother) getName() string {
	return m.Namefull
}
func (m *brother) getName() string {
	return m.Namefull
}
func (m *father) getName() string {
	return m.Name
}

type father struct {
	Name   string
	Old    int
	cookie string
}

func (d *brother) cook() {
	fmt.Printf("Lovely breakfast is done :%s\n", d.Breakfast)
}
func (d *father) cook() {
	fmt.Printf("Lovely breakfast is done :%s\n", d.cookie)
}

func test_breakfast(i cooker) {
	switch i.(type) {
	case *brother:
		i.cook()
	case *father:
		i.cook()
	default:
		fmt.Printf("непонятный персонаж\n")
	}
}

func main() {
	slice := []cooker{&brother{"12312", 11, "123"},
		&father{"12312", 11, "123"},
		&mother{"12312", 11, "123"}}
	mapa := make(map[int]cooker)
	for i := 0; i < 3; i++ {
		mapa[i] = slice[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("test breakfast from %#v\n", mapa[i].getName())
		test_breakfast(mapa[i])
		//fmt.Printf("\n")
	}
}
