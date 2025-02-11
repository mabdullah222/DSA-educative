package questions

func reverseKElementInQueue(que *Queue, k int) {
	//Implement your solution here
	tempq := new(Queue)
	temps := new(Stack)
	for i := 0; i < k; i++ {
		value := que.Dequeue().(int)
		temps.Push(value)
	}
	for !que.IsEmpty() {
		tempq.Enqueue(que.Dequeue())
	}
	for !temps.IsEmpty() {
		que.Enqueue(temps.Pop())
	}
	for !tempq.IsEmpty() {
		que.Enqueue(tempq.Dequeue())
	}
}
