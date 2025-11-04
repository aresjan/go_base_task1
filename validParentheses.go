package main

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if v == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if v == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
			return true
		}
	}
	return len(stack) == 0
}
func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
	println(isValid("]"))
}
