package sortingquestions

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func Partition012(arr []int, size int) {
	left := 0
	right := size - 1
	i := 0
	for i <= right {
		if arr[i] == 0 {
			swap(arr, i, left)
			i += 1
			left += 1
		} else if arr[i] == 2 {
			swap(arr, i, right)
			right -= 1
		} else {
			i += 1
		}
	}
}
