package implementation

import "fmt"

type CircularQueue struct {
	head *Node
	tail *Node
	size int
}

func (q *CircularQueue) Size() int {
	return q.size
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.head.data
}

func (q *CircularQueue) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
	q.tail.next = q.head

}

func (q *CircularQueue) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1, false
	}
	value := q.head.data
	q.head = q.head.next
	q.tail.next = q.head
	q.size--
	return value, true
}

func (q *CircularQueue) Print() {
	temp := q.head
	for temp != q.tail {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Print(temp.data, " ")
	fmt.Println()
}
