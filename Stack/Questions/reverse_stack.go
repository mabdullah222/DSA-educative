package Questions

func reverseStack(stk *Stack) {
	//Implement your solution here
	tempArr := []int{}
	for !stk.IsEmpty() {
		tempArr = append(tempArr, stk.Pop().(int))
	}
	for _, value := range tempArr {
		stk.Push(value)
	}
}
