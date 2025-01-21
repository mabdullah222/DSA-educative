package questions

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := (high + low) / 2
	if data[mid] == value {
		return true
	} else if data[mid] > value {
		return BinarySearchRecursiveUtil(data, low, mid-1, value)
	} else {
		return BinarySearchRecursiveUtil(data, mid+1, high, value)
	}

}
