package questions

import "fmt"

func (t *Tree) PrintDataInRange(min int, max int) {
	printDataInRange(t.root, min, max)
}

func printDataInRange(root *Node, min int, max int) {
	if root == nil {
		return
	}
	//Implement your solution here
	if root.value >= min && root.value <= max {
		printDataInRange(root.left, min, max)
		fmt.Print(root.value, " ")
		printDataInRange(root.right, min, max)
	} else if root.value < min {
		printDataInRange(root.right, min, max)
	} else if root.value > max {
		printDataInRange(root.left, min, max)
	}

	//Implement your solution here
}
