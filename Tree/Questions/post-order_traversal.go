package questions

import "fmt"

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
}

func printPostOrder(n *Node) {

	//Implement your solution here
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value, " ")

	//Implement your solution here

}
