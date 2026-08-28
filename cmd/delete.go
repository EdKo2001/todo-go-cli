package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a todo item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseTodoID(args[0])
		if err != nil {
			return err
		}

		todos, err := loadTodos()
		if err != nil {
			return err
		}

		todo, err := todos.Delete(id)
		if err != nil {
			return err
		}
		if err := saveTodos(todos); err != nil {
			return err
		}

		cmd.Printf("Deleted todo %d: %s\n", todo.ID, todo.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
