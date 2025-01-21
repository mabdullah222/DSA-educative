package Questions

func longestContBalParen(str string, size int) int {
	stk := new(Stack)
	stk.Push(-1)
	length := 0
	for i := 0; i < size; i++ {
		if str[i] == '(' {
			stk.Push(i)
		} else {
			stk.Pop()
			if stk.Length() != 0 {
				if length < i-stk.Top().(int) {
					length = i - stk.Top().(int)
				}
			} else {
				stk.Push(i)
			}
		}
	}
	return length
}
