package todo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile() *os.File {
	file, err := os.Open("data/todos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return file
}

func ReadData(todos *[]Todo) error {
	file := OpenFile();
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Couldn't read Data")
		return err
	}

	for _, todo := range data[1:] {
		id, err := strconv.Atoi(todo[0])
		if err != nil {
			fmt.Println("Couldn't Read data")
			return err
		}

		var completed bool
		temp, _ := strconv.Atoi(todo[2]);
		if temp == 1 {
			completed = true
		} else {
			completed = false
		}

		todo := Todo{id: id, text: todo[1], completed: completed}

		*todos = append(*todos, todo)
	}
	
	return nil;
}

