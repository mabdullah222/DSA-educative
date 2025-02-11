package questions

import "Queue/implementation"

func reverseQueue(que *implementation.Queue) {
	//Implement your solution here
	if que.IsEmpty() {
		return
	}
	temp := que.Dequeue()
	reverseQueue(que)
	que.Enqueue(temp)
}
