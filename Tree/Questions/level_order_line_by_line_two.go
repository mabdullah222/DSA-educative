package questions

import "fmt"

func (t *Tree) PrintLevelOrderLineByLine() {
	que1 := new(Queue)
	que2 := new(Queue)
	var temp *Node
	if t.root != nil {
		que1.Enqueue(t.root)
	}
	for !que1.IsEmpty() || !que2.IsEmpty() {
		for !que1.IsEmpty() {
			temp = que1.Dequeue().(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que2.Enqueue(temp.left)
			}
			if temp.right != nil {
				que2.Enqueue(temp.right)
			}
		}
		fmt.Print("; ")
		for !que2.IsEmpty() {
			temp = que2.Dequeue().(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Enqueue(temp.left)
			}
			if temp.right != nil {
				que1.Enqueue(temp.right)
			}
		}
		fmt.Print("; ")
	}
}
