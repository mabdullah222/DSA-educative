package questions

func MaxMinArr(arr []int, size int) {
	aux := make([]int, size)
	copy(aux, arr)
	start := 0
	stop := size - 1
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			arr[i] = aux[stop]
			stop -= 1
		} else {
			arr[i] = aux[start]
			start += 1
		}
	}
}
