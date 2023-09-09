package main

import "fmt"

var todos []string

func main() {

	var choice int

	fmt.Println("Choose option:")
	fmt.Println("1. View Todos")
	fmt.Println("2. Add Todo")
	fmt.Println("3. Delete Todo")
	fmt.Println("4. Exit")
	fmt.Println("Enter choice: ")
	fmt.Scan(&choice)
	fmt.Println(choice)

	switch choice {
	case 1:
		view()
	case 2:
		add()
	case 3:
		del()
	case 4:
		//
	default:
		fmt.Println("Enter correct option")
	}
}

func view() {

}

func add() {

}

func del() {

}
