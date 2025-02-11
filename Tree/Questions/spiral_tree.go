package questions

import "fmt"

func (t *Tree) PrintSpiralTree() {
	stk1 := new(Stack)
	stk2 := new(Stack)
	root := t.root
	var temp *Node

	if root != nil {
		stk1.Push(root)
	}
	for !stk1.IsEmpty() || !stk2.IsEmpty() {
		for !stk1.IsEmpty() {
			temp = stk1.Pop()
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				stk2.Push(temp.left)
			}
			if temp.right != nil {
				stk2.Push(temp.right)
			}
		}
		fmt.Print("; ")
		for !stk2.IsEmpty() {
			temp = stk2.Pop()
			fmt.Print(temp.value, " ")
			if temp.right != nil {
				stk1.Push(temp.right)
			}
			if temp.left != nil {
				stk1.Push(temp.left)
			}
		}
		fmt.Print("; ")
	}
}
