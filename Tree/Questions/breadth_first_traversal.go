package questions

import (
	"fmt"
)

func (t *Tree) PrintBreadthFirst() {

	//Implement your solution here
	que := new(Queue)
	que.Enqueue(t.root)
	for !que.IsEmpty() {
		tempInterface := que.Dequeue()
		temp, _ := tempInterface.(*Node)
		if temp.left != nil {
			que.Enqueue(temp.left)
		}
		if temp.right != nil {
			que.Enqueue(temp.right)
		}
		fmt.Print(temp.value, " ")
	}
}
