package cmd

import "github.com/spf13/cobra"

var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Mark a todo item as completed",
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

		todo, err := todos.Complete(id)
		if err != nil {
			return err
		}
		if err := saveTodos(todos); err != nil {
			return err
		}

		cmd.Printf("Completed todo %d: %s\n", todo.ID, todo.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
