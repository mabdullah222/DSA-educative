package SortingAlgos

func CountSort(arr []int) []int {
	min := 1000
	max := -1
	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	new_arr := make([]int, max+1)
	for _, value := range arr {
		new_arr[value] += 1
	}
	sorted_arr := []int{}
	for index, value := range new_arr {
		for i := 0; i < value; i++ {
			sorted_arr = append(sorted_arr, index)
		}
	}
	return sorted_arr
}
