package questions

func sumArr(data []int) int {

	total := 0

	//Implement your solution here
	for i := 0; i < len(data); i++ {
		total += data[i]
	}
	return total
}
