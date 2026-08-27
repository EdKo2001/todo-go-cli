package main

import (
	"fmt"
	"strconv"
	"strings"
)

const version = "1.0.0"

func Execute(args []string, todos *Todos) (bool, error) {
	if len(args) == 0 {
		printHelp()
		return false, fmt.Errorf("requires at least one command")
	}

	command := args[0]
	commandArgs := args[1:]

	switch command {
	case "add":
		return runAdd(commandArgs, todos)
	case "list":
		if len(commandArgs) != 0 {
			return false, fmt.Errorf("list does not accept arguments")
		}
		todos.Print()
		return false, nil
	case "done", "toggle":
		return runDone(commandArgs, todos)
	case "edit":
		return runEdit(commandArgs, todos)
	case "delete":
		return runDelete(commandArgs, todos)
	case "clear":
		if len(commandArgs) != 0 {
			return false, fmt.Errorf("clear does not accept arguments")
		}
		*todos = Todos{}
		fmt.Println("Cleared all todos")
		return true, nil
	case "help", "-h", "--help":
		printHelp()
		return false, nil
	case "version", "-v", "--version":
		fmt.Print(version)
		return false, nil
	default:
		return false, fmt.Errorf("unknown command %q for \"todo\"", command)
	}
}

func runAdd(args []string, todos *Todos) (bool, error) {
	var due string
	var titleParts []string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--due":
			if i+1 >= len(args) {
				return false, fmt.Errorf("--due requires a date")
			}
			due = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "-") {
				return false, fmt.Errorf("unknown add flag %q", args[i])
			}
			titleParts = append(titleParts, args[i])
		}
	}

	if len(titleParts) == 0 {
		return false, fmt.Errorf("add requires a task title")
	}

	todo := todos.Add(strings.Join(titleParts, " "), due)
	fmt.Printf("Added todo %d: %s\n", todo.ID, todo.Title)
	return true, nil
}

func runDone(args []string, todos *Todos) (bool, error) {
	index, err := todoIndex(args, todos, "done")
	if err != nil {
		return false, err
	}
	if err := todos.Toggle(index); err != nil {
		return false, err
	}
	fmt.Printf("Completed todo %d\n", (*todos)[index].ID)
	return true, nil
}

func runEdit(args []string, todos *Todos) (bool, error) {
	if len(args) < 2 {
		return false, fmt.Errorf("usage: todo edit <id> <new-title>")
	}

	index, err := findTodoIndex(args[0], todos)
	if err != nil {
		return false, err
	}

	(*todos)[index].Title = strings.Join(args[1:], " ")
	fmt.Printf("Edited todo %d\n", (*todos)[index].ID)
	return true, nil
}

func runDelete(args []string, todos *Todos) (bool, error) {
	index, err := todoIndex(args, todos, "delete")
	if err != nil {
		return false, err
	}

	id := (*todos)[index].ID
	if err := todos.Delete(index); err != nil {
		return false, err
	}
	fmt.Printf("Deleted todo %d\n", id)
	return true, nil
}

func todoIndex(args []string, todos *Todos, command string) (int, error) {
	if len(args) != 1 {
		return -1, fmt.Errorf("usage: todo %s <id>", command)
	}
	return findTodoIndex(args[0], todos)
}

func findTodoIndex(value string, todos *Todos) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil || id < 1 {
		return -1, fmt.Errorf("invalid todo ID %q", value)
	}

	for index := range *todos {
		if (*todos)[index].ID == id {
			return index, nil
		}
	}

	return -1, fmt.Errorf("todo %d not found", id)
}

func printHelp() {
	fmt.Println(`Todo Go - a simple command-line task manager.

Usage:
  todo <command> [arguments] [flags]

Available Commands:
  add <title>             Add a new task
  list                    List tasks
  done <id>               Mark a task as completed
  edit <id> <new-title>   Edit a task
  delete <id>             Delete a task
  clear                   Delete all tasks
  help                    Show help
  version                 Show version

Flags:
  --due <date>            Set a due date when adding a task
  -h, --help              Show help
  -v, --version           Show version

Examples:
  todo add "Learn Go"
  todo add "Finish CLI" --due 2026-08-30
  todo list
  todo done 1`)
}
