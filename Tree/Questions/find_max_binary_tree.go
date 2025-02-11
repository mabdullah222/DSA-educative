package questions

import "math"

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func max(val1 int, val2 int, val3 int) int {
	if val1 > val2 && val1 > val3 {
		return val1
	}
	if val2 > val3 && val2 > val1 {
		return val2
	}
	return val3
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}
	return max(curr.value, findMaxBT(curr.left), findMaxBT(curr.right))
}
