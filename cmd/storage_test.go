package cmd

import (
	"io"
	"path/filepath"
	"testing"
)

func TestAddCommandSavesTodo(t *testing.T) {
	originalFileName := todoFileName
	originalDue := due
	todoFileName = filepath.Join(t.TempDir(), "todos.json")
	due = "2026-09-01"
	addCmd.SetOut(io.Discard)
	t.Cleanup(func() {
		todoFileName = originalFileName
		due = originalDue
		addCmd.SetOut(nil)
	})

	if err := addCmd.RunE(addCmd, []string{"Learn Cobra"}); err != nil {
		t.Fatalf("run add command: %v", err)
	}

	loaded, err := loadTodos()
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
