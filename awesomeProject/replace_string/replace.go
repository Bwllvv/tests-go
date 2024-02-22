package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"
const refStringTwo = "lamb lamb lamb lamb"

func main() {
	out := strings.Replace(refString, " a ", "---", 1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)

	//вывод количества символов а не байт
	s := "asdwsadsdawdawdaw214141_)))#_@$_#$@_#$)@#*$@&#$dллаыладывдаыдуаыДДЖАЖВА?["
	fmt.Println(len([]rune(s)))
	fmt.Println(len(s))
}
