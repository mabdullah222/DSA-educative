package questions

func GCD(m int, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}
