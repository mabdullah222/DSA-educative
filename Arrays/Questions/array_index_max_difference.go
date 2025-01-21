package questions

func ArrayIndexMaxDiff(arr []int, size int) int {
	maxDiff := -1
	j := 0
	for i := 0; i < size; i++ {
		j = size - 1
		for j > i {
			if arr[j] > arr[i] {
				if maxDiff < (j - i) {
					maxDiff = (j - i)
				}

				break
			}
			j -= 1
		}

	}
	return maxDiff
}
