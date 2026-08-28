package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <id> <new-title>",
	Short: "Edit a todo item",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseTodoID(args[0])
		if err != nil {
			return err
		}

		title := strings.TrimSpace(strings.Join(args[1:], " "))
		if title == "" {
			return errors.New("new title cannot be empty")
		}

		todos, err := loadTodos()
		if err != nil {
			return err
		}

		todo, err := todos.Edit(id, title)
		if err != nil {
			return err
		}
		if err := saveTodos(todos); err != nil {
			return err
		}

		cmd.Printf("Edited todo %d: %s\n", todo.ID, todo.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
