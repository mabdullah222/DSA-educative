package Questions

import (
	"fmt"
)

type Stack struct {
	s []interface{}
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *Stack) IsEmpty() bool {
	//Implement your solution here
	return len((*s).s) == 0
	//Kindly make changes according to your needs
}

// length() function returns the number of elements in the stack.
func (s *Stack) Length() int {
	//Implement your solution here
	return len((*s).s) //Kindly make changes according to your needs
}

// The print function will print the elements of the array.
func (s *Stack) Print() {
	//Implement your solution here
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()

}

// push() function append value to the data.
func (s *Stack) Push(value interface{}) {
	//Implement your solution here
	s.s = append(s.s, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *Stack) Pop() interface{} {
	//Implement your solution here
	if s.IsEmpty() {
		return -1
	}
	value := s.s[len(s.s)-1]
	s.s = s.s[0 : len(s.s)-1]
	return value //Kindly make changes according to your needs
}

/*
top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it).
*/
func (s *Stack) Top() interface{} {
	//Implement your solution here
	if s.IsEmpty() {
		return -1
	}
	return s.s[len(s.s)-1] //Kindly make changes according to your needs
}
