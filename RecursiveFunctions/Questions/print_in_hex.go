package questions

import "fmt"

func printInt(number int) {
	conversion := "0123456789ABCDEF"
	base := 16
	digit := number % base
	number = number / base
	if number != 0 {
		printInt(number)
	}
	fmt.Print(string(conversion[digit]))
}
