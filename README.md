# Todo Go CLI

A simple command-line todo application built with Go.

Todo Go is a personal project for practicing core Go concepts while building a useful CLI application.

## Features

- Add tasks with optional due dates
- List all saved tasks in a table
- Mark tasks as completed
- Edit and delete tasks
- Clear all tasks
- Store tasks locally in JSON
- Help and version commands

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

Display all saved tasks:

```bash
todo list
```

Example output:

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
No todos so far. Use `todo add <title>` to create one.
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
Todo Go - a simple command-line task manager.

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
  todo done 1
```

### Version

```bash
todo version
todo -v
todo --version
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
├── command/
│   ├── command.go
│   ├── main.go
│   ├── storage.go
│   └── todo.go
├── flag-practice/
│   ├── README.md
│   └── main.go
└── go.mod
```

The current manual command parser is a standalone program in `command/`.
The repository root is reserved for the future Cobra version.

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
go build -o todo ./command
./todo help
```

During development:

```bash
go run ./command list
```

## Project Goals

Todo Go is designed for learning and practicing:

- Structs and methods
- Pointers and slices
- Command-line arguments and flags
- File I/O and JSON serialization
- Error handling

## Status

Todo Go is currently being developed as a Go learning project.

## License

MIT
