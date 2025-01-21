package questions

func indexArray(arr []int, size int) {
	//Implement your solution here
	status := true
	for status {
		status = false
		for i := 0; i < len(arr); i++ {
			if arr[i] != -1 {
				if arr[i] != i {
					status = true
					arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
				}
			}
		}
	}

	//Note: Please make changes in the given array
}
