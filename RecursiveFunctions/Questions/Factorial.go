package questions

func Factorial(i int) int {
	// Implement your solution here
	if i == 1 {
		return 1
	}
	return i * Factorial(i-1)
}
