package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	Due       string
	CreateAt  time.Time
	// why we do need a pointer here
	// We use a pointer for the CompletedAt field because it allows us to represent the absence of a completion time. If a todo is not completed, the CompletedAt field will be nil, indicating that there is no completion time. If we used a non-pointer type (like time.Time), we would have to use a zero value (e.g., time.Time{}) to represent an uncompleted todo, which could be less clear and less efficient in terms of memory usage.
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) Add(title string, due string) Todo {
	todo := Todo{
		// how does it work
		// The ID is set to len(*t) + 1 to ensure that each new todo gets a unique identifier. The length of the slice *t represents the number of existing todos, so by adding 1, we assign the next available ID to the new todo. This approach assumes that todos are not deleted or that IDs are not reused, which may not be ideal in all cases. In a more robust implementation, you might want to track the highest ID used or use a different method for generating unique IDs.
		ID:        len(*t) + 1,
		Title:     title,
		Completed: false,
		Due:       due,
		CreateAt:  time.Now(),
	}

	*t = append(*t, todo)

	return todo
}

// why is doesn't fire error
// The Add method does not return an error because it is designed to always succeed in adding a new todo to the list. In this implementation, there are no conditions that would cause the addition of a todo to fail (e.g., invalid input, storage issues, etc.). Therefore, the method simply returns the newly created Todo struct without any error handling. If you wanted to add error handling (for example, checking for duplicate titles or validating the due date), you could modify the method to return an error as well.
func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		// ErrorF vs errors.New("index out of range")
		// fmt.Errorf allows you to create formatted error messages, similar to how fmt.Sprintf works for strings. It lets you include dynamic values in the error message. In contrast, errors.New simply creates a new error with a static message. In this case, since we don't need to include any dynamic values in the error message, we could use either fmt.Errorf("index out of range") or errors.New("index out of range"). Both would work fine, but using fmt.Errorf is more flexible if you ever want to include additional context in the future.
		return fmt.Errorf("index out of range")
	}
	return nil
}

func (t *Todos) Delete(index int) error {
	// why do we need to create such var nad not use only t
	// We create a new variable todo to hold the value of *t (the slice of todos) because we want to work with a copy of the slice rather than modifying the original slice directly. This allows us to perform operations like appending or slicing without affecting the original slice until we explicitly assign the modified slice back to *t. It helps in maintaining clarity and avoiding unintended side effects when manipulating slices in Go.
	todo := *t

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*t = append(todo[:index], todo[index+1:]...) // what ... does this do?

	return nil
}

func (t *Todos) Toggle(index int) error {
	todo := *t // why do we need to create such var nad not use only t

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := todo[index].Completed

	if !isCompleted {
		todo[index].Completed = !isCompleted
		now := time.Now()
		todo[index].CompletedAt = &now
	}

	return nil
}

func (t *Todos) Print() {
	if len(*t) == 0 {
		fmt.Println("No todos so far. Use `todo add <title>` to create one.")
		return
	}

	rows := [][]string{
		{"ID", "STATUS", "TITLE", "DUE", "CREATED", "COMPLETED"},
	}

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
