# Todo Go CLI

A simple command-line todo application built with Go.

Todo Go is a personal project for practicing core Go concepts while building a useful CLI application.

## Features

- Add tasks
- Add optional due dates
- List active tasks
- List completed tasks
- List all tasks
- Mark tasks as completed
- Edit existing tasks
- Delete tasks
- Clear all tasks
- Persistent local storage
- Help command
- Version command

## Usage

```bash
todo <command> [arguments] [flags]
```

## Commands

### Add a Task

Add a new task:

```bash
todo add "Learn Go"
```

Add a task with a due date:

```bash
todo add "Finish Todo Go" --due 2026-08-30
```

### List Tasks

Show active tasks:

```bash
todo list
```

Show all tasks:

```bash
todo list --all
```

Show only completed tasks:

```bash
todo list --completed
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

Display the help menu using any of the following:

```bash
todo help
todo -h
todo --help
```

Output:

```text
Todo Go — a simple command-line task manager.

Usage:
  todo <command> [arguments] [flags]

Commands:
  add <title>             Add a new task
  list                    List active tasks
  done <id>               Mark a task as completed
  edit <id> <new-title>   Edit an existing task
  delete <id>             Delete a task
  clear                   Clear all tasks
  help                    Show help
  version                 Show version

Flags:
  add:
    --due <date>           Set a due date (YYYY-MM-DD)

  list:
    --all                  Show all tasks
    --completed            Show completed tasks

Global:
  -h, --help               Show help
  -v, --version            Show version

Examples:
  todo add "Learn Go"
  todo add "Finish CLI" --due 2026-08-30
  todo list
  todo list --completed
  todo list --all
  todo done 2
  todo edit 2 "Finish Todo Go CLI"
  todo delete 2
```

### Version

```bash
todo version
todo -v
todo --version
```

Example output:

```text
Todo Go v1.0.0
```

## Task Model

```go
type Todo struct {
    ID        int
    Title     string
    Completed bool
    Due       string
}
```

## Project Structure

```text
todo-go-cli/
├── go.mod
├── main.go
├── todo.go
├── storage.go
└── README.md
```

## Storage

Tasks are stored locally so they remain available after the CLI exits.
The application serializes task data to JSON and loads it again when the application starts.

Example:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": false,
    "due": ""
  },
  {
    "id": 2,
    "title": "Finish Todo Go",
    "completed": false,
    "due": "2026-08-30"
  }
]
```

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd todo-go-cli
```

Build the application:

```bash
go build -o todo
```

Run it:

```bash
./todo help
```

During development:

```bash
go run . list
```

## Project Goals

Todo Go is designed as a hands-on project for learning and practicing Go concepts, including:

- Structs
- Methods
- Interfaces
- Pointers
- Slices and maps
- Packages
- Command-line arguments
- Flags
- File I/O
- JSON serialization
- Error handling
- Testing

## Status

Todo Go is currently being developed as a Go learning project.

## License

MIT
