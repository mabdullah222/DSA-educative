package questions

func CreateBinarySearchTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinarySearchTreeUtil(arr, 0, size-1)
	return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	curr := &Node{value: arr[mid]}

	curr.left = createBinarySearchTreeUtil(arr, start, mid-1)
	curr.right = createBinarySearchTreeUtil(arr, mid+1, end)

	return curr
}
