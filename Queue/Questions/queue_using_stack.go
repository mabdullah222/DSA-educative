package questions

type QueueUsingStack struct {
	stk1 Stack
}

func (que *QueueUsingStack) Add(value int) {
	que.stk1.Push(value)
}

func (que *QueueUsingStack) Remove() int {
	var temp Stack
	for !que.stk1.IsEmpty() {
		temp.Push(que.stk1.Pop())
	}
	value := temp.Pop()
	for !temp.IsEmpty() {
		que.stk1.Push(temp.Pop())
	}
	return value
}

func (que *QueueUsingStack) Length() int {
	return que.stk1.Length()
}

func (que *QueueUsingStack) IsEmpty() bool {
	return que.stk1.IsEmpty()
}
