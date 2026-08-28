package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var todoFileName = "todos.json"

func loadTodos() (Todos, error) {
	fileData, err := os.ReadFile(todoFileName)
	if errors.Is(err, os.ErrNotExist) {
		return Todos{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read todos: %w", err)
	}

	var todos Todos
	if err := json.Unmarshal(fileData, &todos); err != nil {
		return nil, fmt.Errorf("decode todos: %w", err)
	}

	return todos, nil
}

func saveTodos(todos Todos) error {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("encode todos: %w", err)
	}

	if err := os.WriteFile(todoFileName, fileData, 0644); err != nil {
		return fmt.Errorf("write todos: %w", err)
	}

	return nil
}
