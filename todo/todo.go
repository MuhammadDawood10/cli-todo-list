package todo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Todo struct {
	id        int
	text      string
	completed bool
}

var todoCounter = 1

var reader = bufio.NewReader(os.Stdin)

func GetTodoStatus(isCompleted bool) string {
	if isCompleted {
		return "Done"
	} else {
		return "Not Done"
	}
}

func PrintTodos(todos *[]Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintf(w, "id\tTask\tStatus \n")
	for _, todo := range *todos {
		fmt.Fprintf(w, "%v\t%v\t%v\n", todo.id, todo.text, GetTodoStatus(todo.completed))
	}

	w.Flush()
}

func AddTodo(todos *[]Todo) {
	fmt.Print("Enter text: ")

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	text = strings.TrimSuffix(text, "\n")

	var todo Todo
	todo.id = todoCounter
	todo.completed = false
	todo.text = text

	todoCounter++

	*todos = append(*todos, todo)
}

func DeleteTodo(todos *[]Todo) {
	fmt.Print("Enter number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	id, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	for i, todo := range *todos {
		if todo.id == id {
			*todos = append((*todos)[:i], (*todos)[:i+1]...)
			return
		}
	}

	fmt.Println("Todo not found")
}

func ToggleTodo(todos *[]Todo) {
	fmt.Print("Enter number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	id, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}
	
	for _, todo := range *todos {
		if todo.id == id {
			(*todos)[id - 1].completed = !(*todos)[id - 1].completed;
			return
		}
	}

}
