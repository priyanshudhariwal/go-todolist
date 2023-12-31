package main

import (
	"fmt"
	"os"
)

var todos []string

func main() {

	var choice int

	for {
		fmt.Println("\nChoose option:")
		fmt.Println("1. View Todos")
		fmt.Println("2. Add Todo")
		fmt.Println("3. Delete Todo")
		fmt.Println("4. Exit")
		fmt.Printf("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			view()
		case 2:
			add()
		case 3:
			del()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Enter correct option")
		}
	}
}

func view() {
	fmt.Println("\nCurrent Todos:")
	for i, v := range todos {
		fmt.Printf("%d. %s\n", i, v)
	}
}

func add() {
	var todo string
	fmt.Printf("\nEnter Todo: ")
	fmt.Scanln(&todo)
	todos = append(todos, todo)
	fmt.Println("Successfully added.")
}

func del() {

}
