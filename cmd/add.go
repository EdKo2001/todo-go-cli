package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var due string

var addCmd = &cobra.Command{
	Use:   "add <title>",
	Short: "Add a new todo item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.TrimSpace(args[0])
		if title == "" {
			return errors.New("todo title cannot be empty")
		}

		todos, err := loadTodos()
		if err != nil {
			return err
		}

		todo := todos.Add(title, due)
		if err := saveTodos(todos); err != nil {
			return err
		}

		cmd.Printf("Added todo %d: %s\n", todo.ID, todo.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&due, "due", "", "Set a due date")
}
