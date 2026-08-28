package cmd

import "github.com/spf13/cobra"

var listAll bool
var listCompleted bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		todos, err := loadTodos()
		if err != nil {
			return err
		}

		switch {
		case listAll:
			printTodos(cmd.OutOrStdout(), todos)
		case listCompleted:
			printTodos(cmd.OutOrStdout(), todos.Completed())
		default:
			printTodos(cmd.OutOrStdout(), todos.Active())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listAll, "all", false, "Show all todo items")
	listCmd.Flags().BoolVar(&listCompleted, "completed", false, "Show completed todo items")
	listCmd.MarkFlagsMutuallyExclusive("all", "completed")
}
