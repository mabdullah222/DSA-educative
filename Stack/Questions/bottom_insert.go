package Questions

func bottomInsert(stk *Stack, element int) {
	//Implement your solution here
	var temp int
	if stk.IsEmpty() {
		stk.Push(element)
		return
	}
	temp = stk.Pop().(int)
	bottomInsert(stk, element)
	stk.Push(temp)
}
