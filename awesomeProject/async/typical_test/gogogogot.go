package main

import "fmt"

func main() {
	fmt.Println(handle())
}
func handle() error {

	return &errorr{text: "error"}
}

type error interface {
	error1()
}
type errorr struct {
	text string
}

func (e *errorr) error1() {

}
