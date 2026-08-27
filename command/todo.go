package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

type Todo struct {
	ID          int
	Title       string
	Completed   bool
	Due         string
	CreateAt    time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) Add(title string, due string) Todo {
	todo := Todo{
		ID:        len(*t) + 1,
		Title:     title,
		Completed: false,
		Due:       due,
		CreateAt:  time.Now(),
	}

	*t = append(*t, todo)
	return todo
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("index out of range")
	}
	return nil
}

func (t *Todos) Delete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *Todos) Toggle(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}

	if !(*t)[index].Completed {
		(*t)[index].Completed = true
		now := time.Now()
		(*t)[index].CompletedAt = &now
	}

	return nil
}

func (t *Todos) Print() {
	if len(*t) == 0 {
		fmt.Println("No todos so far. Use `todo add <title>` to create one.")
		return
	}

	rows := [][]string{{"ID", "STATUS", "TITLE", "DUE", "CREATED", "COMPLETED"}}

	for _, todo := range *t {
		status := "Pending"
		if todo.Completed {
			status = "Done"
		}

		due := todo.Due
		if due == "" {
			due = "-"
		}

		completedAt := "-"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("2006-01-02 15:04")
		}

		rows = append(rows, []string{
			fmt.Sprintf("%d", todo.ID),
			status,
			todo.Title,
			due,
			todo.CreateAt.Format("2006-01-02 15:04"),
			completedAt,
		})
	}

	printTable(rows)
}

func printTable(rows [][]string) {
	widths := make([]int, len(rows[0]))
	for _, row := range rows {
		for column, value := range row {
			widths[column] = max(widths[column], utf8.RuneCountInString(value))
		}
	}

	printBorder(widths)
	printTableRow(rows[0], widths)
	printBorder(widths)
	for _, row := range rows[1:] {
		printTableRow(row, widths)
	}
	printBorder(widths)
}

func printBorder(widths []int) {
	fmt.Print("+")
	for _, width := range widths {
		fmt.Print(strings.Repeat("-", width+2), "+")
	}
	fmt.Println()
}

func printTableRow(row []string, widths []int) {
	fmt.Print("|")
	for column, value := range row {
		padding := widths[column] - utf8.RuneCountInString(value)
		fmt.Printf(" %s%s |", value, strings.Repeat(" ", padding))
	}
	fmt.Println()
}
