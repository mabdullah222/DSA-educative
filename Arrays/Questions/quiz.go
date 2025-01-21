package questions

// max function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxPathSum function to calculate the maximum path sum
func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {
	sum1, sum2 := 0, 0
	result := 0
	i, j := 0, 0

	// Traverse both arrays
	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			sum1 += arr1[i]
			i++
		} else if arr1[i] > arr2[j] {
			sum2 += arr2[j]
			j++
		} else {
			// Add the maximum sum plus the common element
			result += max(sum1, sum2) + arr1[i]
			sum1, sum2 = 0, 0
			i++
			j++
		}
	}

	// Add remaining elements of arr1
	for i < size1 {
		sum1 += arr1[i]
		i++
	}

	// Add remaining elements of arr2
	for j < size2 {
		sum2 += arr2[j]
		j++
	}

	// Add the maximum of the remaining sums
	result += max(sum1, sum2)

	return result
}
