package questions

func RotateArray(data []int, k int) {
	//Implement your solution here
	for i := 0; i < k; i++ {
		p1 := 0
		p2 := 1
		for p1 < len(data)-1 {
			temp := data[p2]
			data[p2] = data[p1]
			data[p1] = temp
			p1 += 1
			p2 += 1
		}
	}
	//You must make changes in the given array
}
