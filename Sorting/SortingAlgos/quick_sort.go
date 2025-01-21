package SortingAlgos

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func QuickSort(arr []int) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1)
}
func quickSortUtil(arr []int, lower int, upper int) {
	if lower >= upper {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)
	quickSortUtil(arr, start, upper-1)
	quickSortUtil(arr, upper+1, stop)
}
