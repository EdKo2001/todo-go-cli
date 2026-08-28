# Todo Go CLI

A simple command-line todo application built with Go and Cobra.

Todo Go is a personal project for practicing core Go concepts while building a useful CLI application.

## Features

- Add tasks with optional due dates
- List all saved tasks in a table
- Mark tasks as completed
- Edit and delete tasks
- Clear all tasks
- Store tasks locally in JSON
- Help and version commands
- Automatically generated command help and shell completion

## Usage

```bash
todo <command> [arguments] [flags]
```

## Commands

### Add a Task

```bash
todo add "Learn Go"
todo add "Finish Todo Go" --due 2026-08-30
```

The due date is stored as entered. The CLI does not currently validate its format.

### List Tasks

Display active tasks:

```bash
todo list
```

Display all tasks or only completed tasks:

```bash
todo list --all
todo list --completed
```

Example `todo list --all` output:

```text
+----+---------+------------------+------------+------------------+------------------+
| ID | STATUS  | TITLE            | DUE        | CREATED          | COMPLETED        |
+----+---------+------------------+------------+------------------+------------------+
| 1  | Pending | Learn Go         | -          | 2026-08-25 09:00 | -                |
| 2  | Pending | Finish Todo Go   | 2026-08-30 | 2026-08-25 09:05 | -                |
| 3  | Done    | Learn interfaces | -          | 2026-08-25 09:10 | 2026-08-25 10:30 |
+----+---------+------------------+------------+------------------+------------------+
```

When there are no tasks, the command prints:

```text
No todos found.
```

### Complete a Task

```bash
todo done 2
```

### Edit a Task

```bash
todo edit 2 "Finish Todo Go CLI"
```

### Delete a Task

```bash
todo delete 2
```

### Clear Tasks

```bash
todo clear
```

### Help

```bash
todo help
todo -h
todo --help
```

Output:

```text
A simple command-line todo application

Usage:
  todo [command]

Available Commands:
  add         Add a new todo item
  clear       Clear all todo items
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a todo item
  done        Mark a todo item as completed
  edit        Edit a todo item
  help        Help about any command
  list        List todo items
  version     Show the current version

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.
```

### Version

```bash
todo version
```

Output:

```text
1.0.0
```

## Task Model

```go
type Todo struct {
    ID          int
    Title       string
    Completed   bool
    Due         string
    CreateAt    time.Time
    CompletedAt *time.Time
}
```

`CompletedAt` is `nil` until the task is completed.

## Project Structure

```text
todo-go-cli/
├── .gitignore
├── LICENSE
├── README.md
├── cmd/
│   ├── add.go
│   ├── clear.go
│   ├── commands_test.go
│   ├── delete.go
│   ├── done.go
│   ├── edit.go
│   ├── id.go
│   ├── list.go
│   ├── root.go
│   ├── storage.go
│   ├── storage_test.go
│   ├── table.go
│   ├── todo.go
│   ├── todo_test.go
│   └── version.go
├── command/
│   ├── README.md
│   ├── command.go
│   ├── main.go
│   ├── storage.go
│   └── todo.go
├── flag-practice/
│   ├── README.md
│   └── main.go
├── go.mod
├── go.sum
└── main.go
```

The Cobra implementation starts in `main.go`, with its commands, model, table,
and storage helpers separated under `cmd/`. The original manual parser remains
available as a standalone learning example in `command/`.

## Storage

Tasks are loaded from and saved to `todos.json` in the current directory. The file is created after a command changes the task list.

Example:

```json
[
  {
    "ID": 1,
    "Title": "Learn Go",
    "Completed": false,
    "Due": "",
    "CreateAt": "2026-08-25T09:00:00-05:00",
    "CompletedAt": null
  },
  {
    "ID": 2,
    "Title": "Finish Todo Go",
    "Completed": true,
    "Due": "2026-08-30",
    "CreateAt": "2026-08-25T09:05:00-05:00",
    "CompletedAt": "2026-08-25T10:30:00-05:00"
  }
]
```

## Installation

Clone the repository:

```bash
git clone https://github.com/EdKo2001/todo-go-cli.git
cd todo-go-cli
```

Build and run the application:

```bash
go build -o todo .
./todo help
```

During development:

```bash
go run . list
```

Run the archived manual implementation with:

```bash
go run ./command list
```

## Project Goals

Todo Go is designed for learning and practicing:

- Structs and methods
- Pointers and slices
- Command-line arguments and flags
- Cobra commands and command-specific flags
- File I/O and JSON serialization
- Error handling

## Status

Todo Go is currently being developed as a Go learning project.

## License

MIT
