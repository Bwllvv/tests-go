package main

import "fmt"

func main() {
	stinga := "Let's take LeetCode contest"
	s := []rune(stinga)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 32 || i == len(s)-1 {
			r := start
			k := 0
			if i == len(s)-1 {
				k = i
			} else {
				k = i - 1
			}
			for j := k; j >= r; j-- {
				temp := s[j]
				s[j] = s[r]
				s[r] = temp
				r++
			}
			start = i + 1
		}
	}
	fmt.Println(string(s))
}
