package questions

import (
	"fmt"
)

func (t *Tree) PrintDepthFirst() {
	stk := new(Stack)
	stk.Push(t.root)

	for !stk.IsEmpty() {
		temp := stk.Pop()
		if temp.right != nil {
			stk.Push(temp.right)
		}
		if temp.left != nil {
			stk.Push(temp.left)
		}

		fmt.Print(temp.value, " ")
	}
}
