package main

import (
	"fmt"
)

func main() {

	fmt.Println("start of the project")
	fmt.Print("Please Enter \n 1) For infix Expressions Conversion To NFA\n 2) for postFix Expressions Conversion To NFA \n 3) To Exit\n")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		fmt.Println("Option", input, "Was Entered !!!")
	case 2:
		fmt.Println("Option", input, "Was Entered !!!")
	default:
		fmt.Println("Enter on of the following options")
	}

}
