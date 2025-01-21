package questions

func SmallestPositiveMissingNumber(arr []int, size int) int {
	//Implement your solution here
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 0 && arr[i] <= size {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
	if arr[0] != 1 {
		return 1
	}
	for i := 0; i < len(arr); i++ {
		if (i+1) < len(arr) && arr[i+1]-arr[i] != 1 {
			return arr[i] + 1
		}
	}
	return -1 //Return -1 if missing number not found
}
