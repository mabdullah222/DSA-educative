package questions

import "math"

func (t *Tree) CeilBST(val int) int {

	ceil := math.MinInt32
	temp := t.root
	//Implement your solution here
	for temp != nil {
		if temp.value <= val {
			temp = temp.right
		} else {
			ceil = temp.value
			temp = temp.left
		}
	}
	if ceil < 0 {
		return -1
	}
	return ceil
}
