package main

import (
	"fmt"

	util "./src/util"
)

func main() {
	//This is true unless the user wants to exit
	toExit := true

	//for loop keeps going until the user enters 3 to exit
	for toExit {
		//ask the user which option they like
		fmt.Println("\nstart of the project")
		fmt.Print("Please Enter \n 1) For infix Expressions Conversion To NFA\n 2) For postFix Expressions Conversion To NFA \n 3) To Exit\n")
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			//show which option was choosen
			fmt.Println("Option", input, "Was Entered !!!")
			//promt the user to enter the infix string
			fmt.Print("Enter infix expression: ")
			//use the readfrominput method to get the data from the user
			infixString, err := util.ReadFromInput()
			//error handle the data
			if err != nil {
				fmt.Println("Error when scanning input:", err.Error()) /*  */
				return
			}
			//print the infix the user entered
			fmt.Println("infix", infixString)
			// then send the string to intopost which changes it from infix to postfix
			//(a.(b|d))* used asa example
			newPost := util.IntoPost(infixString)
			fmt.Println("postfix notation:", newPost)

			fmt.Print("Enter a string to test if it matches the nfa: ")
			userTest, err := util.ReadFromInput()

			if err != nil {
				fmt.Println("Error when scanning input:", err.Error()) /*  */
				return
			}

			fmt.Println("Does the string", userTest, " match ?", util.Pomatch(newPost, userTest))

		case 2:
			//show which option was choosen
			fmt.Println("Option", input, "Was Entered !!!")
			//promt the user to enter the infix string
			fmt.Print("Enter postfix expression: ")
			//use the readfrominput method to get the data from the user
			infixString, err := util.ReadFromInput()
			//error handle the data
			if err != nil {
				fmt.Println("Error when scanning input:", err.Error()) /*  */
				return
			}

			fmt.Print("Enter a string to test if it matches the nfa: ")
			userTest, err := util.ReadFromInput()

			if err != nil {
				fmt.Println("Error when scanning input:", err.Error()) /*  */
				return
			}

			fmt.Println("Does the string", userTest, " match ?", util.Pomatch(infixString, userTest))
		default:
			toExit = false
			fmt.Println("Enter one of the above options")
		}
	}
}
