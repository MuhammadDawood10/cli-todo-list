package main

import (
	"bufio"
	"cli-todo-list/todo"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Todo List")

	todos := []todo.Todo{}

	todo.ReadData(&todos)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		todo.PrintTodos(&todos)
		fmt.Println("Commands: add, delete, complete, uncomplete, exit ")

		fmt.Print("Enter command: ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			todo.AddTodo(&todos)
		case "delete":
			todo.DeleteTodo(&todos)
		case "exit":
			return
		case "complete":
			todo.ToggleTodo(&todos)
		case "uncomplete":
			todo.ToggleTodo(&todos)
		default:
			log.Fatal("Invalid Command!")
		}

	}

}
