package main

import "fmt"

func main() {
	fir := "leetcode"
	sec := "leet"
	fmt.Println(strStr(fir, sec))
}
func strStr(heystack string, needle string) int {
	count := 0
	for i := 0; i < len(heystack); i++ {
		for len(heystack)-1 >= count+i && len(needle)-1 >= count && heystack[i+count] == needle[count] {
			count++
		}
		if count == len(needle) {
			return i
		}
		count = 0
	}
	return -1
}
