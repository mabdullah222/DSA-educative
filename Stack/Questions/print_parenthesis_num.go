package Questions

import (
	"fmt"
	"strconv"
)

func printParenthesisNumber(expn string, size int) {
	//Uncomment the ch variable and use it iterate through the string
	count := 1
	stack := new(Stack)
	output := ""

	for _, ch := range expn {
		if ch == '(' {
			stack.Push(count)
			output += strconv.Itoa(count) // Convert number to string
			count++
		} else if ch == ')' {
			value := stack.Pop().(int)
			output += strconv.Itoa(value) // Convert number to string
		}
	}

	//Implement your solution here

	fmt.Println(output)
}
