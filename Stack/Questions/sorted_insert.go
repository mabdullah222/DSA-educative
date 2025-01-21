package Questions

func sortedInsert(stk *Stack, element int) {
	//Implement your solution here
	if stk.IsEmpty() {
		stk.Push(element)
	}
	new_stk := []int{}
	for stk.Top().(int) > element {
		new_stk = append(new_stk, stk.Pop().(int))
	}
	stk.Push(element)
	for i := len(new_stk) - 1; i >= 0; i-- {
		stk.Push(new_stk[i])
	}

}
