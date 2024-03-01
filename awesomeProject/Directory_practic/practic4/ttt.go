package main

func main() {
	stri := "{{}}[()]"
	println(isValid(stri))
}
func isValid(stri string) bool {
	var stack []rune
	for _, value := range stri {
		switch value {
		case '(', '[', '{':
			stack = append(stack, value)
		case ')', ']', '}':
			top := stack[len(stack)-1]
			if top == '(' && value != ')' || top == '[' && value != ']' || top == '{' && value != '}' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
