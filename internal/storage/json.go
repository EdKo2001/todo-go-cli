package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/EdKo2001/todo-go-cli/internal/todo"
)

type JSONStore struct {
	FileName string
}

func (store JSONStore) Load() (todo.Todos, error) {
	fileData, err := os.ReadFile(store.FileName)
	if errors.Is(err, os.ErrNotExist) {
		return todo.Todos{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read todos: %w", err)
	}

	var todos todo.Todos
	if err := json.Unmarshal(fileData, &todos); err != nil {
		return nil, fmt.Errorf("decode todos: %w", err)
	}

	return todos, nil
}

func (store JSONStore) Save(todos todo.Todos) error {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("encode todos: %w", err)
	}

	if err := os.WriteFile(store.FileName, fileData, 0644); err != nil {
		return fmt.Errorf("write todos: %w", err)
	}

	return nil
}
