package questions

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}
func WaveArray(arr []int) {
	size := len(arr)
	//Implement your solution here
	for i := 1; i < size; i += 2 {
		if (i-1) >= 0 && arr[i] > arr[i-1] {
			swap(arr, i, i-1)
		}
		if (i+1) < size && arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
	}
	//Update the same array
}
