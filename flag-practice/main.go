package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0"

type options struct {
	add     string
	due     string
	list    bool
	done    int
	delete  int
	clear   bool
	help    bool
	version bool
}

func main() {
	options := parseFlags()

	if err := execute(options); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func parseFlags() options {
	var options options

	flag.StringVar(&options.add, "add", "", "Add a todo")
	flag.StringVar(&options.due, "due", "", "Set a due date")
	flag.BoolVar(&options.list, "list", false, "List todos")
	flag.IntVar(&options.done, "done", -1, "Complete a todo by ID")
	flag.IntVar(&options.delete, "delete", -1, "Delete a todo by ID")
	flag.BoolVar(&options.clear, "clear", false, "Clear all todos")
	flag.BoolVar(&options.help, "help", false, "Show help")
	flag.BoolVar(&options.help, "h", false, "Show help")
	flag.BoolVar(&options.version, "version", false, "Show version")
	flag.BoolVar(&options.version, "v", false, "Show version")

	flag.Parse()

	return options
}

func execute(options options) error {
	if options.help {
		printHelp()
		return nil
	}

	if options.version {
		fmt.Printf("Todo Flags Practice v%s\n", version)
		return nil
	}

	switch {
	case options.add != "":
		fmt.Printf("ADD title=%q due=%q\n", options.add, options.due)
	case options.list:
		fmt.Println("LIST")
	case options.done != -1:
		if options.done < 1 {
			return fmt.Errorf("-done requires an ID greater than zero")
		}
		fmt.Printf("DONE id=%d\n", options.done)
	case options.delete != -1:
		if options.delete < 1 {
			return fmt.Errorf("-delete requires an ID greater than zero")
		}
		fmt.Printf("DELETE id=%d\n", options.delete)
	case options.clear:
		fmt.Println("CLEAR")
	default:
		printHelp()
	}

	return nil
}

func printHelp() {
	fmt.Println(`Todo Flags Practice

Usage:
  go run ./flag-practice [flags]

Flags:
  -add <title>       Add a todo
  -due <date>        Set a due date with -add
  -list              List todos
  -done <id>         Complete a todo
  -delete <id>       Delete a todo
  -clear             Clear all todos
  -h, -help          Show help
  -v, -version       Show version

Examples:
  go run ./flag-practice -add "Learn flags"
  go run ./flag-practice -add "Ship project" -due 2026-09-01
  go run ./flag-practice -list
  go run ./flag-practice -done 2`)
}
