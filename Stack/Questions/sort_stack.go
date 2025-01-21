package Questions

func sortStack(stk *Stack) {
	var temp int
	if stk.Length() > 0 {
		temp = stk.Pop().(int)
		sortStack(stk)
		sortedInsert(stk, temp)
	}
}
