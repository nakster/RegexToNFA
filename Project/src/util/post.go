package util

//this func turns the infix string to postfix
func IntoPost(infix string) string {
	//this a map that is going to map speial characters like . * | into numbers or integers
	specials := map[rune]int{'*': 10, '|': 9, '+': 7, '?': 7, '.': 7}
	//this stores the postfix string
	postfix := []rune{}
	//we put the infix on this
	stack := []rune{}

	//range converts each cahracter of the string into a rune
	for _, r := range infix {
		switch {
		case r == '(':
			//if there is a ( then this gets appended to the stack
			stack = append(stack, r)
		case r == ')':
			//if you see closed bracket keep popping things of the stack and
			//keep putting then into postfix
			//until you see a open bracket
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//discard the open bracket
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			//if the length of the stack is greater than zero
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default:
			postfix = append(postfix, r)
		}
	}
	//error handling
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	//finally return the string
	return string(postfix)
}
