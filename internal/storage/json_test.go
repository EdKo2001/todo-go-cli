package storage

import (
	"path/filepath"
	"testing"

	"github.com/EdKo2001/todo-go-cli/internal/todo"
)

func TestJSONStoreSavesAndLoadsTodos(t *testing.T) {
	store := JSONStore{FileName: filepath.Join(t.TempDir(), "todos.json")}
	todos := todo.Todos{}
	todos.Add("Learn Cobra", "2026-09-01")

	if err := store.Save(todos); err != nil {
		t.Fatalf("save todos: %v", err)
	}

	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("load todos: %v", err)
	}
	if len(loaded) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(loaded))
	}
	if loaded[0].ID != 1 || loaded[0].Title != "Learn Cobra" || loaded[0].Due != "2026-09-01" {
		t.Fatalf("unexpected saved todo: %#v", loaded[0])
	}
}

func TestJSONStoreLoadsEmptyListWhenFileDoesNotExist(t *testing.T) {
	store := JSONStore{FileName: filepath.Join(t.TempDir(), "missing.json")}

	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("load missing file: %v", err)
	}
	if len(loaded) != 0 {
		t.Fatalf("expected no todos, got %d", len(loaded))
	}
}
