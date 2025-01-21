package Questions

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

//Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	return s.size //Return 0 if stack is empty
}

/* IsEmpty() function will return true if size of the linked list is
equal to zero or false in all other cases. */
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

/*First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list. */
func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false // Return 0, false if stack is empty
	}
	return s.head.value, true
}

//Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	newNode := &Node{value: value, next: s.head}
	s.head = newNode
	s.size++
}

/*In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it. */
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false // Return 0, false if stack is empty
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	current := s.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
