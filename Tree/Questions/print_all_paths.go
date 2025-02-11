package questions

import "fmt"

func (t *Tree) PrintAllPath() {
	stk := new(Stack)
	printAllPath(t.root, stk)
}

func printAllPath(curr *Node, stk *Stack) {
	stk.Push(curr.value)

	if curr.left == nil && curr.left == nil {
		stk.Print()
		fmt.Print("; ")
		stk.Pop()
		return
	}

	if curr.right != nil {
		printAllPath(curr.right, stk)
	}
	if curr.left != nil {
		printAllPath(curr.left, stk)
	}
	stk.Pop()
}
