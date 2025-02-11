package questions

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {

	//Implement your solution here
	if start >= size {
		return nil
	}
	root := &Node{arr[start], nil, nil}
	left_child := 2*start + 1
	right_child := 2*start + 2

	root.left = levelOrderBinaryTree(arr, left_child, size)
	root.right = levelOrderBinaryTree(arr, right_child, size)

	return root
}
