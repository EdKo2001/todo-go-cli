package cmd

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/EdKo2001/todo-go-cli/internal/todo"
)

func printTodos(output io.Writer, todos todo.Todos) {
	if len(todos) == 0 {
		fmt.Fprintln(output, "No todos found.")
		return
	}

	rows := [][]string{{"ID", "STATUS", "TITLE", "DUE", "CREATED", "COMPLETED"}}
	for _, todo := range todos {
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

	widths := make([]int, len(rows[0]))
	for _, row := range rows {
		for column, value := range row {
			widths[column] = max(widths[column], utf8.RuneCountInString(value))
		}
	}

	printBorder(output, widths)
	printRow(output, rows[0], widths)
	printBorder(output, widths)
	for _, row := range rows[1:] {
		printRow(output, row, widths)
	}
	printBorder(output, widths)
}

func printBorder(output io.Writer, widths []int) {
	fmt.Fprint(output, "+")
	for _, width := range widths {
		fmt.Fprint(output, strings.Repeat("-", width+2), "+")
	}
	fmt.Fprintln(output)
}

func printRow(output io.Writer, row []string, widths []int) {
	fmt.Fprint(output, "|")
	for column, value := range row {
		padding := widths[column] - utf8.RuneCountInString(value)
		fmt.Fprintf(output, " %s%s |", value, strings.Repeat(" ", padding))
	}
	fmt.Fprintln(output)
}
