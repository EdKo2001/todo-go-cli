package cmd

import (
	"os"

	"github.com/EdKo2001/todo-go-cli/internal/storage"
	"github.com/spf13/cobra"
)

var todoStore = storage.JSONStore{FileName: "todos.json"}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple command-line todo application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
