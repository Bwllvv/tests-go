package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "domovoyнуклеотид"
	for _, element := range str {
		fmt.Printf("%c", element)

	}
	fmt.Println(len(str))

	sum := 0
	for t := 0; t < len(str); t++ {
		fmt.Println(fmt.Sprintf("%c", str[t]), sum)
		sum++
	}

	str1 := "Hello, 世界!"
	for _, char := range str1 {
		if utf8.RuneLen(char) > 2 {
			fmt.Printf("%c\n", char)
		}
	}

}
