package Questions

import "fmt"

func maxDepthParenthesis(expn string, size int) int {
	//Implement your solution here

	//Uncomment the ch variable and use it to move in an expression array
	stack := new(Stack)
	maxDepth := 0
	currentDepth := 0

	for i := 0; i < len(expn); i++ {
		ch := expn[i]
		fmt.Println(ch)
		switch ch {
		case '(', '{', '[':
			stack.Push(ch)
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		case ')', '}', ']':
			if !stack.IsEmpty() {
				_ = stack.Pop() // Ignore error since we're checking IsEmpty first
				currentDepth--
			}
		}
	}

	return maxDepth
}
