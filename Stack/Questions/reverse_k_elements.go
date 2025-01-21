package Questions

func reverseKElementInStack(stk *Stack, k int) {
	que := []int{}
	i := 0
	for stk.Length() > 0 && i < k {
		que = append(que, stk.Pop().(int))
		i++
	}
	for _, value := range que {
		stk.Push(value)
	}
}
