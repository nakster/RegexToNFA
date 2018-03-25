package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//ask the user which option they like
	fmt.Println("start of the project")
	fmt.Print("Please Enter \n 1) For infix Expressions Conversion To NFA\n 2) for postFix Expressions Conversion To NFA \n 3) To Exit\n")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		//show which option was choosen
		fmt.Println("Option", input, "Was Entered !!!")
		//promt the user to enter the infix string
		fmt.Print("Enter infix expression: ")
		//use the readfrominput method to get the data from the user
		infixString, err := ReadFromInput()
		//error handle the data
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}
		//print the infix the user entered
		fmt.Println("infix", infixString)
		// then send the string to intopost which changes it from infix to postfix
		fmt.Println("postfix notation:", intoPost(infixString))

	case 2:
		fmt.Println("Option", input, "Was Entered !!!")
	default:
		fmt.Println("Enter on of the following options")
	}

}

//reads user input
func ReadFromInput() (string, error) {
	//read
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

//this func turns the infix string to postfix
func intoPost(infix string) string {
	//this a map that is going to map speial characters like . * | into numbers or integers
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
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
