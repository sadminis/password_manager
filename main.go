package main

import "fmt"

func main() {
	var menu string
	var input string
	menu = "Enter your option: \n1. Store Username/Password\n2. Search Username/Password"

	fmt.Println(menu)
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Print("Entered 1 \n")
	case "2":
		fmt.Print("Entered 2 \n")
	default:
		fmt.Print("Invalid input \n")
	}
}
