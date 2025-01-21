package Questions

// Get operator precedence
func getPrecedence(op rune) int {
	if op == '/' || op == '*' {
		return 4
	} else if op == '+' || op == '-' {
		return 2
	}
	return 0
}

// Convert Infix expression to Postfix
func InfixToPostfix(expn string) string {
	output := "" // Output string
	stack := &Stack{}

	for _, ch := range expn {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
			// Operand: Append directly to output
			output += string(ch)
		} else if ch == '(' {
			// Opening parenthesis: Push onto the stack
			stack.Push(ch)
		} else if ch == ')' {
			// Closing parenthesis: Pop until '(' is found
			for !stack.IsEmpty() && stack.Top() != '(' {
				output += string(stack.Pop().(rune))
			}
			// Pop the '(' from the stack
			stack.Pop()
		} else {
			// Operator: Pop operators of higher or equal precedence
			for !stack.IsEmpty() && getPrecedence(stack.Top().(rune)) >= getPrecedence(ch) {
				output += string(stack.Pop().(rune))
			}
			// Push the current operator onto the stack
			stack.Push(ch)
		}
	}

	// Pop any remaining operators from the stack
	for !stack.IsEmpty() {
		output += string(stack.Pop().(rune))
	}

	return output
}
