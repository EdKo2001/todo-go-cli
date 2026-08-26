package main

import (
	"fmt"
	"os"
)

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")

	if err := storage.Load(&todos); err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}

	changed, err := RunCommand(os.Args[1:], &todos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if changed {
		if err := storage.Save(todos); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving todos: %v\n", err)
			os.Exit(1)
		}
	}
}
