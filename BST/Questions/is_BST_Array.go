package questions

func isBSTArray(preorder []int, size int) bool {

	//Implement your solution here
	lowerBound := -1
	stack := new(Stack)
	for _, value := range preorder {
		if value < lowerBound {
			return false
		}
		for !stack.IsEmpty() && value > stack.Top() {
			lowerBound = stack.Pop()
		}
		stack.Push(value)
	}
	return true
}
