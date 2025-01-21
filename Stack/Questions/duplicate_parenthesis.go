package Questions

func findDuplicateParenthesis(expn string, size int) bool {
	stk := new(Stack)
	var ch byte
	var count int

	for i := 0; i < size; i++ {
		ch = expn[i]
		if ch == ')' {
			count = 0
			for stk.Length() != 0 && byte(stk.Top().(rune)) != '(' {
				stk.Pop()
				count += 1
			}
			if count <= 1 {
				return true
			}
		} else {
			stk.Push(int(ch))
		}
	}
	return false
}
