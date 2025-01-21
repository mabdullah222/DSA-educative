package Questions

func IsBalancedParenthesis(expn string) bool {
	stk := new(Stack)
	for _, ch := range expn {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val.(rune) != '{' {
				return false
			}
		case ']':
			val := stk.Pop().(rune)
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop().(rune)
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}
