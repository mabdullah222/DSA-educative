package questions

func (t *Tree) Find(value int) bool {
	temp := t.root
	for temp != nil {
		if temp.value == value {
			return true
		} else if temp.value > value {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}

	return false
}
