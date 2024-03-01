package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	stri := "😃﷽.!пр美国人"
	fmt.Println(string(stri[0]), stri[0], len(stri))
	runes := []rune(stri)
	fmt.Println(runes[0])
	fmt.Println(utf8.RuneLen(runes[0]))
	//for _, value := range stri {
	//	fmt.Println(string(value))
	//}
}

//001
//010
//100
//011
//110
//111
//101
//011
