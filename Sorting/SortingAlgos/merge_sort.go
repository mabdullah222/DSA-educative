package SortingAlgos

func MergeSort(arr []int) {
	temp_arr := make([]int, len(arr))
	size := len(arr)
	mergesrt(arr, temp_arr, 0, size-1)
}

func mergesrt(arr []int, tempArr []int, lower int, upper int) {
	if lower >= upper {
		return
	}
	middle := (lower + upper) / 2
	mergesrt(arr, tempArr, lower, middle)
	mergesrt(arr, tempArr, middle+1, upper)
	merge(arr, tempArr, lower, middle, upper)
}

func merge(arr []int, tempArr []int, lower int, middle int, upper int) {
	lowerStart := lower
	lowerEnd := middle
	upperStart := middle + 1
	upperEnd := upper
	count := lower

	for lowerStart <= lowerEnd && upperStart <= upperEnd {
		if arr[lowerStart] < arr[upperStart] {
			tempArr[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArr[count] = arr[upperStart]
			upperStart++
		}
		count++
	}
	for lowerStart <= lowerEnd {
		tempArr[count] = arr[lowerStart]
		lowerStart++
		count++
	}

	for upperStart <= upperEnd {
		tempArr[count] = arr[upperStart]
		upperStart++
		count++
	}
	for i := lower; i <= upper; i++ {
		arr[i] = tempArr[i]
	}

}
