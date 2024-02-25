package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Id     int
	Name   string
	Old    int
	Number string
}

var jsonstr = `{"id":42,"name":"lev","old":19,"number":"89174353673"}`

func main() {
	data := []byte(jsonstr)
	u := &person{}
	json.Unmarshal(data, u)
	fmt.Printf("%#v", u)
	person := person{32131, "levichka", 19, "89174353673"}
	json.Marshal(person)
	fmt.Printf("\n%#v", person)
}
