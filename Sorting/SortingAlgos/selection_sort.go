package SortingAlgos

func SelectionSort(arr []int) {
	i := len(arr) - 1
	for i > 0 {
		max := i
		j := 0
		for j < i {
			if arr[j] > arr[max] {
				max = j
			}
			j++
		}
		arr[i], arr[max] = arr[max], arr[i]
		i--
	}
}
