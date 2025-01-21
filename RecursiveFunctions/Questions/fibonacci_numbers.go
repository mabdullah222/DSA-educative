package questions

func fibonacci(n int) int {
	//Implement your solution here
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
