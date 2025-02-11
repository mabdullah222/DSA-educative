package implementation

import "fmt"

type Node struct {
	data int
	next *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.head.data
}

func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	temp := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return temp.data
}

func (q *QueueLinkedList) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
}
