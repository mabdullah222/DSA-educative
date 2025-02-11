package main

import (
	implementation "Queue/Implementation"
	"fmt"
)

// Testing Code
func main() {
	q := new(implementation.Deque)
	q.AddFront(1)
	q.AddFront(1)
	q.AddBack(2)
	q.AddBack(2)
	q.Print()
	fmt.Println(q.RemoveBack())
	q.Print()
	fmt.Println(q.RemoveFront())
	q.Print()
}
