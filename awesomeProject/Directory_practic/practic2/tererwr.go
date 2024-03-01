package main

import "fmt"

func main() {
	stringg := []string{"ac", "ac", "a", "a"}
	fmt.Println(longestCommonPrefix(stringg))
}
func longestCommonPrefix(strs []string) string {
	str := ""
	lastletter := strs[0]
	count := 0
	if len(strs) == 0 {
		return str
	}
	if len(strs) == 1 {
		return lastletter
	}
	for i := 1; i < 2; i++ {
		if len(lastletter) > len(strs[i]) {
			for j := 0; j < len(strs[i]); j++ {
				if lastletter[j] == strs[i][j] {
					str += string(lastletter[j])
					count++
				} else {
					break
				}
			}
			lastletter = strs[i]
		} else {
			for j := 0; j < len(lastletter); j++ {
				if lastletter[j] == strs[i][j] {
					str += string(lastletter[j])
					count++
				} else {
					break
				}
			}

			lastletter = strs[i]
		}
	}

	str2 := ""
	for i := 2; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return str2
		}
		if len(strs[i]) > len(str) {
			for j := 0; j < len(str); j++ {
				if strs[i][j] == str[j] {
					str2 += string(strs[i][j])
					count++
				} else {
					str = str2
					str2 = ""
					break
				}
			}
			if count == 1 {
				str = str2
			}
			str2 = ""
			count = 0
			//lastletter = strs[i]
		} else {
			for j := 0; j < len(strs[i]); j++ {
				if strs[i][j] == str[j] {
					str2 += string(strs[i][j])
					count++
				} else {
					str = str2
					str2 = ""
					break
				}
			}
			if count == 1 {
				str = str2
			}
			str2 = ""
			count = 0
		}
	}

	return str
}
