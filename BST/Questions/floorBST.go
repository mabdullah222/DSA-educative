package questions

import "math"

func (t *Tree) FloorBST(val int) int {

	floor := math.MaxInt32
	temp := t.root
	//Implement your solution here
	for temp != nil {
		if temp.value < val {
			floor = temp.value
			temp = temp.right
		} else {
			temp = temp.left
		}
	}
	return floor
}
