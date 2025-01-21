package questions

func BinarySearch(data []int, value int) bool {
	start := 0
	end := len(data) - 1

	for start <= end {
		middle := start + (end-start)/2

		if data[middle] == value {
			return true
		}
		if data[middle] < value {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false // Return false if value not found
}
