package main

import "fmt"

type animal interface {
	makeSound()
}

type cat struct {
}
type dog struct {
}

func (c *cat) makeSound() {
	fmt.Println("meow")
}
func (d *dog) makeSound() {
	fmt.Println("wow")
}

func main() {
	var c, d animal = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()
	SayHello(&russian{}, "Дима")
	SayHello(&american{}, "Daniel")
}

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

func (g *russian) greet(name string) string {
	return fmt.Sprintf("привет,%s", name)
}
func (g *american) greet(name string) string {
	return fmt.Sprintf("hello,%s", name)
}

func SayHello(v greeter, name string) {
	fmt.Println(v.greet(name))
}
