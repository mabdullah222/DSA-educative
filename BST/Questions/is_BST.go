package questions

func IsBST(root *Node) bool {
	//Implement your solution here
	if root == nil {
		return true
	}
	status1 := false
	status2 := false
	if root.left == nil || (root.left != nil && root.left.value < root.value) {
		status1 = true
	}
	if root.right == nil || (root.right != nil && root.right.value > root.value) {
		status2 = true
	}

	return status1 && status2 && IsBST(root.left) && IsBST(root.right)
}
