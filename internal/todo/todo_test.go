package todo

import "testing"

func TestTodoOperations(t *testing.T) {
	todos := Todos{}
	first := todos.Add("First", "")
	second := todos.Add("Second", "2026-09-01")

	completed, err := todos.Complete(first.ID)
	if err != nil {
		t.Fatalf("complete todo: %v", err)
	}
	if !completed.Completed || completed.CompletedAt == nil {
		t.Fatal("expected todo to be completed with a completion time")
	}

	edited, err := todos.Edit(second.ID, "Updated second")
	if err != nil {
		t.Fatalf("edit todo: %v", err)
	}
	if edited.Title != "Updated second" {
		t.Fatalf("unexpected edited title: %q", edited.Title)
	}

	if len(todos.Active()) != 1 || len(todos.Completed()) != 1 {
		t.Fatal("expected one active and one completed todo")
	}

	deleted, err := todos.Delete(first.ID)
	if err != nil {
		t.Fatalf("delete todo: %v", err)
	}
	if deleted.ID != first.ID || len(todos) != 1 {
		t.Fatal("todo was not deleted correctly")
	}
}

func TestMissingTodoReturnsError(t *testing.T) {
	todos := Todos{}

	if _, err := todos.Complete(99); err == nil {
		t.Fatal("expected an error for a missing todo")
	}
	if _, err := todos.Edit(99, "Missing"); err == nil {
		t.Fatal("expected an error for a missing todo")
	}
	if _, err := todos.Delete(99); err == nil {
		t.Fatal("expected an error for a missing todo")
	}
}
