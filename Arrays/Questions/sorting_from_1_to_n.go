package questions

func Sort1toN(arr []int, size int) {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 1 {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
}
