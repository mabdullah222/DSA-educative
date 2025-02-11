package questions

import "fmt"

func (t *Tree) PrintLevelOrderLineByLine2() {
	que := new(Queue)
	var temp *Node
	root := t.root
	count := 0
	if root != nil {
		que.Enqueue(root)
	}
	for !que.IsEmpty() {
		count = que.Length()
		for count != 0 {
			temp = que.Dequeue().(*Node)
			if temp.left != nil {
				que.Enqueue(temp.left)
			}
			if temp.right != nil {
				que.Enqueue(temp.right)
			}
			fmt.Print(temp.value, " ")
			count -= 1
		}
		fmt.Print("; ")
	}
}
