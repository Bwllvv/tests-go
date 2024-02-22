package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("appalon")
	var p = "name"
	p = "mary"
	fmt.Println(p)
	const dsda int = 32
	fmt.Println(test())
	fmt.Println(solution("abc", "c"))
	var a [4]string
	word := "dom"
	a[0] = "Ыалами"
	a[2] = "fix"
	sum := 0
	//q := fmt.Sprintf("%d", a[0])
	all := a[0] + word
	for sum < len(a) {
		fmt.Println(a[sum], all, len([]rune(a[0])))
		sum++
	}
	//str := "Hello, World!"
	c := a[0]
	str := "кукущкab"
	c1, _ := utf8.DecodeRuneInString(c)
	fmt.Println("00000", c1)
	fmt.Printf("%c", c1)
	for index, char := range str {
		fmt.Println(fmt.Sprintf("%c", char), index)
	}

}
func test() (ac string, ab int) {
	ac = "hello"
	ab = 32
	return
}

func solution(str, ending string) bool {
	phrase := str
	phrase2 := ending
	fmt.Println(len(ending))
	fmt.Println(phrase[0])

	for i := 0; i > -(len(ending) - 1); i-- {

		if phrase[i] == phrase2[i] {
			return true
		}

	}
	return false
}
