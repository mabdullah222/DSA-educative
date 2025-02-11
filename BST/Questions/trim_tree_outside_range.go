package questions

func (t *Tree) TrimOutsidedataRange(min int, max int) {
	t.root = trimOutsidedataRange(t.root, min, max)
}

func trimOutsidedataRange(curr *Node, min int, max int) *Node {
	//Implement your solution here
	if curr == nil {
		return nil
	}
	if curr.value >= min && curr.value <= max {
		curr.left = trimOutsidedataRange(curr.left, min, max)
		curr.right = trimOutsidedataRange(curr.right, min, max)
	}
	if curr.value < min {
		return trimOutsidedataRange(curr.right, min, max)
	}
	if curr.value > max {
		trimOutsidedataRange(curr.left, min, max)
	}

	return curr
}
