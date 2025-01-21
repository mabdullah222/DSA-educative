package SortingAlgos

func InsertionSort(arr []int) {
	j := 0
	for i := 1; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
