package cmd

import (
	"github.com/EdKo2001/todo-go-cli/internal/todo"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all todo items",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := todoStore.Save(todo.Todos{}); err != nil {
			return err
		}

		cmd.Println("Cleared all todo items")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
