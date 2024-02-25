package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Id     int `json:"user_id,string"`
	Name   string
	Old    int    `json:",omitempty"`
	Number string `json:"-"`
}

func main() {

	person := &person{32131, "levichka", 19, "89174353673"}
	result, _ := json.Marshal(person)
	fmt.Printf("\n%s", result)
}
