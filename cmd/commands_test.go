package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/EdKo2001/todo-go-cli/internal/storage"
)

func TestStoredCommandFlow(t *testing.T) {
	originalStore := todoStore
	originalDue := due
	originalListAll := listAll
	originalListCompleted := listCompleted

	todoStore = storage.JSONStore{FileName: filepath.Join(t.TempDir(), "todos.json")}
	due = ""
	listAll = false
	listCompleted = false

	var output bytes.Buffer
	addCmd.SetOut(&output)
	doneCmd.SetOut(&output)
	editCmd.SetOut(&output)
	deleteCmd.SetOut(&output)
	listCmd.SetOut(&output)
	clearCmd.SetOut(&output)

	t.Cleanup(func() {
		todoStore = originalStore
		due = originalDue
		listAll = originalListAll
		listCompleted = originalListCompleted
		addCmd.SetOut(nil)
		doneCmd.SetOut(nil)
		editCmd.SetOut(nil)
		deleteCmd.SetOut(nil)
		listCmd.SetOut(nil)
		clearCmd.SetOut(nil)
	})

	if err := addCmd.RunE(addCmd, []string{"First"}); err != nil {
		t.Fatalf("add first todo: %v", err)
	}
	due = "2026-09-01"
	if err := addCmd.RunE(addCmd, []string{"Second"}); err != nil {
		t.Fatalf("add second todo: %v", err)
	}
	if err := doneCmd.RunE(doneCmd, []string{"1"}); err != nil {
		t.Fatalf("complete todo: %v", err)
	}
	if err := editCmd.RunE(editCmd, []string{"2", "Updated second"}); err != nil {
		t.Fatalf("edit todo: %v", err)
	}
	if err := deleteCmd.RunE(deleteCmd, []string{"1"}); err != nil {
		t.Fatalf("delete todo: %v", err)
	}

	listAll = true
	output.Reset()
	if err := listCmd.RunE(listCmd, nil); err != nil {
		t.Fatalf("list todos: %v", err)
	}
	if !strings.Contains(output.String(), "Updated second") {
		t.Fatalf("list output does not contain edited todo:\n%s", output.String())
	}

	if err := clearCmd.RunE(clearCmd, nil); err != nil {
		t.Fatalf("clear todos: %v", err)
	}
	loaded, err := todoStore.Load()
	if err != nil {
		t.Fatalf("load cleared todos: %v", err)
	}
	if len(loaded) != 0 {
		t.Fatalf("expected no todos after clear, got %d", len(loaded))
	}
}
