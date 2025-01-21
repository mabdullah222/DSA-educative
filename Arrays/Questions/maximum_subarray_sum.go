package questions

func MaxSubArraySum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max_sum := data[0]
	summ := 0

	for i := 0; i < len(data); i++ {
		summ += data[i]
		if summ < 0 {
			summ = 0
		}
		if summ > max_sum {
			max_sum = summ
		}
	}

	return max_sum
}
