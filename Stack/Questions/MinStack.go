package Questions

type MinStack struct {
	maxSize      int
	mainStack    Stack
	minimumStack Stack
}

// removes and returns value from stack
func (m *MinStack) Pop() int {
	_ = m.minimumStack.Pop()
	top := m.mainStack.Top()
	_ = m.mainStack.Pop()
	return top.(int)
}

// Pushes value into the stack
func (m *MinStack) Push(value int) {
	m.mainStack.Push(value)
	if !m.minimumStack.IsEmpty() && m.minimumStack.Top().(int) < value {
		m.minimumStack.Push(m.minimumStack.Top())
	} else {
		m.minimumStack.Push(value)
	}
}

// returns minimum value in O(1)
func (m *MinStack) Min() int {
	return m.minimumStack.Top().(int)
}
