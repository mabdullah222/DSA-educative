package questions

import "fmt"

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
}

func printPreOrder(n *Node) {

	//Implement your solution here
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)

	//Implement your solution here

}
