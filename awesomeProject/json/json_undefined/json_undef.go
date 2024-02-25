package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `[
	{"id": 17, "username": "iivan", "phone": 0},
	{"id": "17", "address": "none", "company": "Mail.ru"}
]`

func main() {
	data := []byte(jsonStr)
	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("%#v", user1)
}
