package questions

import (
	"fmt"
	"strconv"
)

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func Permutation(data []int, i int, length int) {

	// base condition DO NOT ALTER IT
	if length == i {
		temp := "{"
		for k := 0; k < length; k++ {
			temp += strconv.Itoa(data[k])
			temp += " "
		}
		temp += "} "
		fmt.Print(temp)
		return
	}
	// Write your code here
	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}
